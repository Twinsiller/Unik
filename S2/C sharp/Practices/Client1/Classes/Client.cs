using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;

namespace Clients.Classes
{
    internal class Client
    {
        string _nameClient;

        public Client(string nameClient = "unknown")
        {
            _nameClient = nameClient;
        }

        public void giveFile(string filePath)
        {
            if (!File.Exists(filePath))
            {
                Console.WriteLine("Файл не найден.");
                return;
            }

            string fileName = Path.GetFileName(filePath); // Имя файла для отправки
            byte[] fileNameBytes = Encoding.UTF8.GetBytes(fileName);
            long fileSize = new FileInfo(filePath).Length;

            try
            {
                using (TcpClient client = new TcpClient("127.0.0.1", 12345)) // Подключаемся к серверу
                using (NetworkStream stream = client.GetStream()) // Поток для передачи данных
                using (BinaryWriter writer = new BinaryWriter(stream)) // Для отправки данных
                {
                    Console.WriteLine("Подключено к серверу. Отправка данных...");

                    // 1. Отправляем длину имени файла
                    writer.Write(fileNameBytes.Length);

                    // 2. Отправляем имя файла
                    writer.Write(fileNameBytes);

                    // 3. Отправляем размер файла
                    writer.Write(fileSize);

                    // 4. Отправляем содержимое файла
                    using (FileStream fileStream = new FileStream(filePath, FileMode.Open, FileAccess.Read))
                    {
                        byte[] buffer = new byte[4096];
                        int bytesRead;
                        while ((bytesRead = fileStream.Read(buffer, 0, buffer.Length)) > 0)
                        {
                            stream.Write(buffer, 0, bytesRead);
                        }
                    }

                    Console.WriteLine("Файл успешно отправлен!");
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Ошибка: {ex.Message}");
            }
        }


    }
}
