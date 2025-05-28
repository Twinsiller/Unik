using Pr4_Client.ClassOpt;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace Pr4_Client.Frames._2Body
{
    /// <summary>
    /// Логика взаимодействия для MainBody.xaml
    /// </summary>
    public partial class MainBody : Page
    {
        private ClientEngine _client = new();

        public MainBody()
        {
            InitializeComponent();
            _client.LogAction = msg => OutputBox.Items.Add(msg);
            _ = _client.ConnectAsync("127.0.0.1", 12345);
        }

        private async void Send_Click(object sender, RoutedEventArgs e)
        {
            string guess = GuessInput.Text;
            await _client.SendGuessAsync(guess);
        }
    }
}
