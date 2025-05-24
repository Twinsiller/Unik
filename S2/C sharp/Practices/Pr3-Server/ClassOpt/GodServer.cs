using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Pr3_Server.ClassOpt
{
    class GodServer
    {
        private readonly List<FileReceiverServer> frss;
        public GodServer() {
            frss = new List<FileReceiverServer>();
        }

        public void AddServer(int port)
        {
            foreach (FileReceiverServer frs in frss)
            {
                if (frs._port == port) {
                    Console.WriteLine($"Port already exist: {port}");
                    return;
                }
            }
            frss.Add(new FileReceiverServer(port/*, &Files*/));
            Console.WriteLine($"Port was added: {port}");
        }

        public void DeleteServer(int port)
        {
            foreach (FileReceiverServer frs in frss)
            {
                if (frs._port == port)
                {
                    frss.Remove(frs);
                    Console.WriteLine($"Port was deleted: {port}");
                    return;
                }
            }
            Console.WriteLine($"Port doesnt exist: {port}");
        }
    }
}
            