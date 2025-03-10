using System;
using System.IO;
using System.Net.Sockets;
using System.Text;
using Clients.Classes;

namespace Clients
{
    internal class StartClients
    {
        static void Main(string[] args)
        {
            Client client1 = new Client("Bob");
            string filePath;
            while (true)
            {
                Console.Write("Введите путь к файлу для отправки: ");
                filePath = Console.ReadLine();
                if (filePath == "") break;
                client1.giveFile(filePath);
            }

            Console.Write("Программа завершена");
        }
    }
}
