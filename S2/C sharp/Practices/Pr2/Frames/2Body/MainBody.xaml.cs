using Microsoft.Win32;
using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;
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
using Pr2.Services;

namespace Pr2.Frames._2Body
{
    /// <summary>
    /// Логика взаимодействия для MainBody.xaml
    /// </summary>
    public partial class MainBody : Page
    {
        private readonly FilesAnalisator _analisator = new FilesAnalisator();
        
        public MainBody()
        {
            InitializeComponent();
        }

        private async void SelectFiles_Click(object sender, RoutedEventArgs e)
        {
            var openFileDialog = new OpenFileDialog
            {
                Multiselect = true,
                Filter = "Все файлы (*.*)|*.*"
            };

            if (openFileDialog.ShowDialog() == true)
            {
                try
                {
                    FileList_Name.Items.Clear();
                    FileList_Info.Items.Clear();
                    int countFile = 1;
                    // Показываем имена файлов до обработки
                    foreach (var file in openFileDialog.FileNames)
                    {
                        FileList_Name.Items.Add($"Файл {countFile++}: {System.IO.Path.GetFileName(file)}");
                        FileList_Info.Items.Add($"Обработка: {System.IO.Path.GetFileName(file)}");
                    }
                    TextBlockFilesCount.Text = $"Количество файлов: {countFile - 1}";
                    // Асинхронный анализ
                    var results = await _analisator.AnalyzeFilesAsync(openFileDialog.FileNames);

                    // Обновляем UI с результатами
                    FileList_Info.Items.Clear();
                    foreach (var result in results)
                    {
                        FileList_Info.Items.Add(
                            $"{result._fileName} - Слов: {result._countWords}, Символов: {result._countChars}");
                    }

                    // Итоговая статистика
                    int totalWords = results.Sum(r => r._countWords);
                    int totalChars = results.Sum(r => r._countChars);
                    FileList_Info.Items.Add($"ИТОГО: Слов - {totalWords}, Символов - {totalChars}");
                }
                catch (Exception ex)
                {
                    MessageBox.Show($"Ошибка: {ex.Message}", "Ошибка",
                        MessageBoxButton.OK, MessageBoxImage.Error);
                }
            }
        }
    }
}
