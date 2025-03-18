using ServerWPF_Pr2_.ClassOpt;
using ServerWPF_Pr2_.Pages;
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

namespace ServerWPF_Pr2_
{
    /// <summary>
    /// Логика взаимодействия для MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();

            FrameNav.frmHeaderSubject = ViewFrmHeadSubject;
            ViewFrmHeadSubject.Navigate(new Subject());
            
            FrameNav.frmHeaderPracticeNumber = ViewFrmHeadPracticeNumber;
            ViewFrmHeadPracticeNumber.Navigate(new PracticeNumber());
        }

        private void Button_Click(object sender, RoutedEventArgs e)
        {

        }
    }
}
