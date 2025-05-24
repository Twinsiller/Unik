using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Net;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using Pr3_Server.ClassOpt;
using System.Collections.ObjectModel;
using System.Windows;
using Pr3_Server.Frames._2Body;

namespace Pr3_Server.ClassOpt
{
    class FileReceiverServer
    {
        //public ObservableCollection<FileInfoView> Files { get; set; }

        private readonly string directoryFileName = "ReceivedFiles";
        public int _port { get; }
        private TcpListener _server;
        private CancellationTokenSource _cts;

        public FileReceiverServer(int port)
        {
            _port = port;
            Start();
        }

        public void Start()
        {
            _cts = new CancellationTokenSource();
            Task.Run(() => RunServer(_cts.Token));
        }

        public void Stop()
        {
            _cts?.Cancel();
            _server?.Stop();
            Console.WriteLine($"Сервер на порту {_port} остановлен.");
        }

        private async Task RunServer(CancellationToken token)
        {
            try
            {
                _server = new TcpListener(IPAddress.Loopback, _port);
                _server.Start();
                Console.WriteLine($"Сервер запущен на порту {_port}. Ожидание подключений...");

                while (!token.IsCancellationRequested)
                {
                    if (_server.Pending())
                    {
                        TcpClient client = await _server.AcceptTcpClientAsync();
                        Console.WriteLine("Клиент подключен.");
                        _ = Task.Run(() => HandleReadFileFromPol(client));
                    }
                    else
                    {
                        await Task.Delay(100, token); // Немного подождать, если нет клиентов
                    }
                }
            }
            catch (Exception ex) when (!(ex is OperationCanceledException))
            {
                Console.WriteLine($"Ошибка в сервере: {ex.Message}");
            }
        }

        private void HandleReadFileFromPol(TcpClient client)
        {
            try
            {
                using (client)
                using (NetworkStream stream = client.GetStream())
                using (BinaryReader reader = new BinaryReader(stream))
                {
                    Console.WriteLine("Приём файла...");

                    int fileNameLength = reader.ReadInt32();
                    string fileName = Encoding.UTF8.GetString(reader.ReadBytes(fileNameLength));
                    long fileSize = reader.ReadInt64();

                    Directory.CreateDirectory(directoryFileName);
                    string savePath = Path.Combine(directoryFileName, fileName);

                    // Уникальное имя файла
                    if (File.Exists(savePath))
                    {
                        string ext = Path.GetExtension(fileName);
                        string name = Path.GetFileNameWithoutExtension(fileName);
                        int counter = 1;
                        do
                        {
                            fileName = $"{name}_{counter}{ext}";
                            savePath = Path.Combine(directoryFileName, fileName);
                            counter++;
                        } while (File.Exists(savePath));
                    }

                    using (FileStream fileStream = new FileStream(savePath, FileMode.Create, FileAccess.Write))
                    {
                        byte[] buffer = new byte[4096];
                        long totalRead = 0;
                        int bytesRead;

                        while (totalRead < fileSize && (bytesRead = stream.Read(buffer, 0, buffer.Length)) > 0)
                        {
                            fileStream.Write(buffer, 0, bytesRead);
                            totalRead += bytesRead;
                        }
                    }

                    Console.WriteLine($"Файл '{fileName}' успешно принят!");
                    Application.Current.Dispatcher.Invoke(() =>
                    {
                        var mainBody = FrameNav.conrtollerMainBody.Content as MainBody;

                        if (mainBody != null)
                        {
                            mainBody.AddFile(fileName); 
                        }
                    });

                }
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Ошибка при приёме файла: {ex.Message}");
            }
        }

        ~FileReceiverServer()
        {
            Stop();
        }
    }
}
