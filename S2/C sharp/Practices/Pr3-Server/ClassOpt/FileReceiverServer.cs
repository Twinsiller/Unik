using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Net;
using System.Text;
using System.Threading.Tasks;

namespace Server.Classes
{
    internal class FileReceiverServer
    {
        string directoryFileName = "ReceivedFiles";

        public int _port { get; }

        public FileReceiverServer(int port)
        {
            _port = port;
            Start();
        }

        

        private void HandleReadFileFromPol(TcpClient client)
        {
            try
            {
                Console.WriteLine($"Точно поймал");
                using (client)
                using (NetworkStream stream = client.GetStream())
                using (BinaryReader reader = new BinaryReader(stream))
                {
                    Console.WriteLine("Клиент подключен. Приём файла...");

                    int fileNameLength = reader.ReadInt32(); // Читаем имя файла
                    string fileName = System.Text.Encoding.UTF8.GetString(reader.ReadBytes(fileNameLength));


                    long fileSize = reader.ReadInt64(); // Читаем размер файла


                    string savePath = Path.Combine(directoryFileName, fileName); // Путь к папке, где хранится файл


                    Directory.CreateDirectory(directoryFileName); // Создаем директорию, если она не существует


                    if (File.Exists(savePath)) // Проверяем, существует ли файл с таким именем
                    {
                        // Изменяем имя файла, добавляя индекс к имени (например, file_1.txt)
                        string fileExtension = Path.GetExtension(fileName); // Получаем расширение файла (например, .txt)
                        string fileNameWithoutExtension = Path.GetFileNameWithoutExtension(fileName); // Имя файла без расширения

                        int counter = 1;
                        // Формируем новое имя файла, добавляя индекс, если файл с таким именем уже существует
                        do
                        {
                            // Новый файл с индексом (fileName_1.txt, fileName_2.txt и так далее)
                            fileName = $"{fileNameWithoutExtension}_{counter}{fileExtension}";
                            savePath = Path.Combine(directoryFileName, fileName);
                            counter++;
                        } while (File.Exists(savePath)); // Повторяем, пока не найдем уникальное имя
                    }


                    using (FileStream fileStream = new FileStream(savePath, FileMode.Create, FileAccess.Write)) // Теперь сохраняем файл по уникальному пути
                    {
                        byte[] buffer = new byte[4096];
                        long totalBytesRead = 0;
                        int bytesRead;

                        while (totalBytesRead < fileSize && (bytesRead = stream.Read(buffer, 0, buffer.Length)) > 0)
                        {
                            fileStream.Write(buffer, 0, bytesRead);
                            totalBytesRead += bytesRead;
                        }
                    }


                    Console.WriteLine($"Файл '{fileName}' успешно принят!");
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Ошибка: {ex.Message}");
            }
        }

        private void Start()
        {
            TcpListener server = new TcpListener(IPAddress.Loopback, _port);
            server.Start();
            Console.WriteLine($"Сервер запущен на порту {_port}. Ожидание подключений...");

            while (true)
            {
                Console.WriteLine($"...");
                TcpClient client = server.AcceptTcpClient(); // Ожидание клиента
                Console.WriteLine($"Поймал");
                Task.Run(() => HandleReadFileFromPol(client)); // Запускаем обработку в отдельной задаче
            }
        }

        ~FileReceiverServer()
        {

        }
    }
}
