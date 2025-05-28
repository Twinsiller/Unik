using Pr4_Server.ClassOpt;
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

namespace Pr4_Server.Frames._2Body
{
    /// <summary>
    /// Логика взаимодействия для MainBody.xaml
    /// </summary>
    public partial class MainBody : Page
    {
        private ServerEngine? _server;
        
        public MainBody()
        {
            InitializeComponent();
        }        

        private async void StartServer_Click(object sender, RoutedEventArgs e)
        {
            _server = new ServerEngine(12345);
            _server.LogAction = msg => LogBox.Items.Add(msg);
            await _server.StartAsync();
        }
    }
}
