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
                //Console.Write("Введите путь к файлу для отправки: ");
                //filePath = Console.ReadLine();
                Console.Write("Quit? (y/n): ");
                string ch = Console.ReadLine();
                filePath = "C:\\Users\\boldi\\Desktop\\Something\\Unik\\S2\\C sharp\\Practices\\Client1\\file.txt";
                if (filePath == "" || ch == "y") break;
                client1.giveFile(filePath);
            }

            Console.Write("Программа завершена");
        }
    }
}
