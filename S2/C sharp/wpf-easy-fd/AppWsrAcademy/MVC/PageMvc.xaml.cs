using AppWsrAcademy.MVC.Controller;
using AppWsrAcademy.MVC.HelpConroller;
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

namespace AppWsrAcademy.MVC
{
    /// <summary>
    /// Логика взаимодействия для PageMvc.xaml
    /// </summary>
    public partial class PageMvc : Page
    {
        public PageMvc()
        {
            InitializeComponent();
        }

        private void BtnCheck_Click(object sender, RoutedEventArgs e)
        {
            //ClientInfo.LoginCheck = TxbLogin.Text;
            MessageBox.Show(ControllerLogin.CheckLoginOdb(TxbLogin.Text));
        }
    }
}
