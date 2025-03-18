using AppWsrAcademy.DataFilesApp;
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

namespace AppWsrAcademy.Teacher
{
    /// <summary>
    /// Логика взаимодействия для PageStudentList.xaml
    /// </summary>
    public partial class PageStudentList : Page
    {
        public PageStudentList()
        {
            InitializeComponent();

            CmbSelectGroup.SelectedValuePath = "Id";
            CmbSelectGroup.DisplayMemberPath = "Name";
            CmbSelectGroup.ItemsSource = OdbConnectHelper.entObj.NameGroup.ToList();

            GridList.ItemsSource = OdbConnectHelper.entObj.Student.ToList();
        }

        private void BtnBack_Click(object sender, RoutedEventArgs e)
        {
            FrameApp.frmObj.GoBack();
        }

        private void BtnSearch_Click(object sender, RoutedEventArgs e)
        {
            //int SelectGroup = Convert.ToInt32(CmbSelectGroup.SelectedValue);
            //GridList.ItemsSource = OdbConnectHelper.entObj.Student.Where(x => x.IdNameGroup == SelectGroup).ToList();
            //GridList.SelectedIndex = 0;
        }

        private void BtnProfile_Click(object sender, RoutedEventArgs e)
        {
            FrameApp.frmObj.Navigate(new PageJournalStudent((sender as Button).DataContext as Student));
        }

        private void CmbSelectGroup_SelectionChanged(object sender, SelectionChangedEventArgs e)
        {
            int SelectGroup = Convert.ToInt32(CmbSelectGroup.SelectedValue);
            GridList.ItemsSource = OdbConnectHelper.entObj.Student.Where(x => x.IdNameGroup == SelectGroup).ToList();
            GridList.SelectedIndex = 0;
        }
    }
}
