using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;
using System.Windows;

namespace Pr4_Client.ClassOpt
{
    public class ClientEngine
    {
        private readonly TcpClient _client = new();
        private StreamReader? _reader;
        private StreamWriter? _writer;

        public Action<string>? LogAction;

        public async Task ConnectAsync(string host, int port)
        {
            await _client.ConnectAsync(host, port);
            var stream = _client.GetStream();
            _reader = new StreamReader(stream);
            _writer = new StreamWriter(stream) { AutoFlush = true };
            _ = ListenAsync();
        }

        private async Task ListenAsync()
        {
            while (true)
            {
                string? msg = await _reader!.ReadLineAsync();
                if (msg != null)
                {
                    Application.Current.Dispatcher.Invoke(() => LogAction?.Invoke(msg));
                }
            }
        }

        public async Task SendGuessAsync(string guess)
        {
            if (_writer != null)
                await _writer.WriteLineAsync(guess.ToUpper());
        }
    }

}
