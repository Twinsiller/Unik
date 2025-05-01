using Pr3_Server.ClassOpt;
using Server.Classes;
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

namespace Pr3_Server.Frames._2Body
{
    /// <summary>
    /// Логика взаимодействия для MainBody.xaml
    /// </summary>
    public partial class MainBody : Page
    {
        GodServer gs = new GodServer();
        public MainBody()
        {
            InitializeComponent();
        }

        private async void StartAwaitingFiles(object sender, RoutedEventArgs e)
        {

            gs.AddServer();
        }

    }
}
