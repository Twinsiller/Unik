using System.Net.Sockets;
using System.Net;
using Server.Classes;

namespace Server
{
    internal class Server
    {
        static void Main()
        {
            FileReceiverServer frs = new FileReceiverServer();
            frs.Start();
        }
    }
}
