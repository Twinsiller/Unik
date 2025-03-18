using AppWsrAcademy.DataFilesApp;
using System;
using System.Collections.Generic;
using System.Data.Odbc;
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

namespace AppWsrAcademy.Teacher
{
    /// <summary>
    /// Логика взаимодействия для PageDeleteStudent.xaml
    /// </summary>
    public partial class PageDeleteStudent : Page
    {
        public PageDeleteStudent()
        {
            InitializeComponent();

            CmbSelectGroup.SelectedValuePath = "Id";
            CmbSelectGroup.DisplayMemberPath = "Name";
            CmbSelectGroup.ItemsSource = OdbConnectHelper.entObj.NameGroup.ToList();

        }

        private void CmbSelectGroup_SelectionChanged(object sender, SelectionChangedEventArgs e)
        {
            int SelectGroupStudent = Convert.ToInt32(CmbSelectGroup.SelectedValue);
            GridListStudent.ItemsSource = OdbConnectHelper.entObj.Student.Where(x => x.IdNameGroup == SelectGroupStudent).ToList();
        }

        private void BtnDelete_Click(object sender, RoutedEventArgs e)
        {
            try
            {
                if (GridListStudent.SelectedItems.Count > 0)
                {
                    for (int i=0; i < GridListStudent.SelectedItems.Count; i++) {
                        Student studentObj = GridListStudent.SelectedItems[i] as Student;
                        OdbConnectHelper.entObj.Student.Remove(studentObj);
                    }

                    OdbConnectHelper.entObj.SaveChanges();
                    MessageBox.Show(
                        "Данные о студенте успешно удалены!",
                        "Уведомление",
                        MessageBoxButton.OK,
                        MessageBoxImage.Information
                        );
                    GridListStudent.ItemsSource = OdbConnectHelper.entObj.Student.ToList();
                }
                else {
                    MessageBox.Show(
                        "Данных в таблице нет!",
                        "Уведомление",
                        MessageBoxButton.OK,
                        MessageBoxImage.Information
                        );
                }
            }
            catch (Exception ex) {
                MessageBox.Show(
                    "Критическая работа приложения: " + ex.Message.ToLower(),
                    "Уведомление",
                    MessageBoxButton.OK,
                    MessageBoxImage.Error
                    );
            }
        }
    }
}
