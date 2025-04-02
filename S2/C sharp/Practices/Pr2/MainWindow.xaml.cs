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
using Pr2.ClassOpt;
using Pr2.Frames._1Header;
using Pr2.Frames._2Body;
using Pr2.Frames._3Footer;

namespace Pr2
{
    /// <summary>
    /// Логика взаимодействия для MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();

            FrameNav.conrtollerMyName = NamePlace;
            NamePlace.Navigate(new MyName());

            FrameNav.conrtollerMainBody = MainBodyPlace;
            MainBodyPlace.Navigate(new MainBody());
            
            FrameNav.conrtollerSubject = SubjectPlace;
            SubjectPlace.Navigate(new Subject());
        }
    }
}
