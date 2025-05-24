using Pr3_Server.ClassOpt;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
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
        ObservableCollection<FileInfoView> Files { get; set; } = new ObservableCollection<FileInfoView>();
        public MainBody()
        {
            InitializeComponent();
            FilesDataGrid.ItemsSource = Files;
        }

        public void AddFile(string fileName)
        {
            Files.Add(new FileInfoView
            {
                FileName = fileName
            });
        }

        private void StartAwaitingFiles(object sender, RoutedEventArgs e)
        {
            try
            {
                gs.AddServer(int.Parse(serverPort.Text.Trim()));
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Ошибка (Нужны цЫфры): {ex.Message}");
            }
           
        }

        private void StopAwaitingFiles(object sender, RoutedEventArgs e)
        {
            try
            {
                gs.DeleteServer(int.Parse(serverPort.Text.Trim()));
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Ошибка (Нужны цЫфры): {ex.Message}");
            }
        }
    }
}
