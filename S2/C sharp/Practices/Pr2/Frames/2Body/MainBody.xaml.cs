using Microsoft.Win32;
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

namespace Pr2.Frames._2Body
{
    /// <summary>
    /// Логика взаимодействия для MainBody.xaml
    /// </summary>
    public partial class MainBody : Page
    {
        public MainBody()
        {
            InitializeComponent();
        }

        private void SelectFile_Click(object sender, RoutedEventArgs e)
        {
            OpenFileDialog openFileDialog = new OpenFileDialog() { Multiselect = true };
            openFileDialog.Filter = "Все файлы (*.*)|*.*"; // Фильтр форматов
            if (openFileDialog.ShowDialog() == true)
            {
                FileList.Items.Clear(); // Очищаем список
                _results.Clear();

                foreach (var file in openFileDialog.FileNames) // Вывод путей в ListBox
                {
                    FileList.Items.Add(System.IO.Path.GetFileName(file)); // Добавляем только имя файла
                }
                   
            }
            await ProcessFilesAsync(openFileDialog.FileNames); // Запускаем обработку файлов

        }
    }
}
