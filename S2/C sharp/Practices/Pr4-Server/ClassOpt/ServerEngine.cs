using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;
using System.Windows;

namespace Pr4_Server.ClassOpt
{
    public class ServerEngine
    {
        private readonly TcpListener _listener;
        private readonly List<TcpClient> _clients = new();
        private readonly GameLogic _logic = new();
        private readonly int _maxPlayers = 4;
        private readonly int _minPlayers = 2;

        public Action<string>? LogAction; // Логирование в UI

        public ServerEngine(int port)
        {
            _listener = new TcpListener(IPAddress.Any, port);
        }

        public async Task StartAsync()
        {
            _listener.Start();
            Log("Сервер запущен.");
            _logic.GenerateCode();
            Log($"Секретный код: {_logic.SecretCode}");

            while (_clients.Count < _maxPlayers)
            {
                TcpClient client = await _listener.AcceptTcpClientAsync();
                _clients.Add(client);
                Log($"Игрок {_clients.Count} подключён.");
                _ = HandleClientAsync(client, _clients.Count);
            }
        }

        private async Task HandleClientAsync(TcpClient client, int playerId)
        {
            using var stream = client.GetStream();
            using var reader = new StreamReader(stream);
            using var writer = new StreamWriter(stream) { AutoFlush = true };

            await writer.WriteLineAsync("Игра началась! Введите код (4 буквы A-Z)");

            while (true)
            {
                string? guess = await reader.ReadLineAsync();
                if (string.IsNullOrWhiteSpace(guess) || guess.Length != 4) continue;

                var result = _logic.Evaluate(guess.ToUpper());
                await writer.WriteLineAsync($"Черные: {result.black}, Белые: {result.white}");

                Log($"Игрок {playerId} → {guess.ToUpper()} → Черные: {result.black}, Белые: {result.white}");

                if (result.black == 4)
                {
                    await writer.WriteLineAsync("Вы победили!");
                    Log($"Игрок {playerId} победил!");
                    break;
                }
            }
        }

        private void Log(string message)
        {
            Application.Current.Dispatcher.Invoke(() => LogAction?.Invoke(message));
        }
    }

}
