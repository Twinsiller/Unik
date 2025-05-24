using Microsoft.Win32;
using Pr3_Client.ClassOpt;
using System;
using System.Collections.Generic;
using System.IO;
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

namespace Pr3_Client.Frames._2Body
{
    /// <summary>
    /// Логика взаимодействия для MainBody.xaml
    /// </summary>
    public partial class MainBody : Page
    {
        Client client = new Client();
        string path { get; set; }
        public MainBody()
        {
            InitializeComponent();
        }

        private void SelectFile_Click(object sender, RoutedEventArgs e)
        {
            var openFileDialog = new OpenFileDialog
            {
                Multiselect = false,
                Filter = "Все файлы (*.*)|*.*"
            };

            if (openFileDialog.ShowDialog() == true)
            {
                try
                {
                    FileList_Name.Items.Clear();
                    // Показываем имена файлов до обработки
                    path = openFileDialog.FileName; 
                    FileList_Name.Items.Add($"Файл: {System.IO.Path.GetFileName(path)}");
                }
                catch (Exception ex)
                {
                    MessageBox.Show($"Ошибка: {ex.Message}", "Ошибка",
                    MessageBoxButton.OK, MessageBoxImage.Error);
                }
            }
        } 
        
        private void SendFile_Click(object sender, RoutedEventArgs e)
        {
            if (path == "")
            {   
                Console.WriteLine("Нет файла");
                return;
            }
            client.sendFile(path, int.Parse(serverPort.Text.Trim()));
            Console.WriteLine("Файл отправлен");
        }
    }
}
