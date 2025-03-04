using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Net;
using System.Text;
using System.Threading.Tasks;

namespace Server.Classes
{
    class FileReceiverServer
    {
        static string directoryFileName = "ReceivedFiles";
        
        private int _port;
        
        public FileReceiverServer(int port = 12345)
        {
            _port = port;
        }

        private void HandleReadFileFromPol(TcpClient client)
        {
            using (client)
            using (NetworkStream stream = client.GetStream())
            using (BinaryReader reader = new BinaryReader(stream))
            {
                Console.WriteLine("Клиент подключен. Приём файла...");

                // Читаем имя файла
                int fileNameLength = reader.ReadInt32();
                string fileName = System.Text.Encoding.UTF8.GetString(reader.ReadBytes(fileNameLength));

                // Читаем размер файла
                long fileSize = reader.ReadInt64();

                // Путь к папке, где будет сохраняться файл
                string savePath = Path.Combine(directoryFileName, fileName);

                // Создаем директорию, если она не существует
                Directory.CreateDirectory(directoryFileName);

                // Проверяем, существует ли файл с таким именем
                if (File.Exists(savePath))
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

                // Теперь сохраняем файл по уникальному пути
                using (FileStream fileStream = new FileStream(savePath, FileMode.Create, FileAccess.Write))
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

        public void Start()
        {
            TcpListener server = new TcpListener(IPAddress.Loopback, _port);
            server.Start();
            Console.WriteLine($"Сервер запущен на порту {_port}. Ожидание подключений...");

            while (true)
            {
                TcpClient client = server.AcceptTcpClient(); // Ожидание клиента
                Task.Run(() => HandleReadFileFromPol(client)); // Запускаем обработку в отдельной задаче
            }
        }

        static void Main()
        {
            FileReceiverServer frs = new FileReceiverServer();
            frs.Start();
        }
    }
}
