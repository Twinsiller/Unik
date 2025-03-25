using ServerWPF_Pr2_.ClassOpt;
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

namespace ServerWPF_Pr2_.MainPages.Header
{
    /// <summary>
    /// Логика взаимодействия для Head.xaml
    /// </summary>
    public partial class Head : Page
    {
        public Head()
        {
            InitializeComponent();

            FrameNav.frmHeaderSubject = ViewFrmHeadSubject;
            ViewFrmHeadSubject.Navigate(new Subject());

            FrameNav.frmHeaderPracticeNumber = ViewFrmHeadPracticeNumber;
            ViewFrmHeadPracticeNumber.Navigate(new PracticeNumber());
        }
    }
}
