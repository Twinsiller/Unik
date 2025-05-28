using Pr4_Server.ClassOpt;
using Pr4_Server.Frames._1Header;
using Pr4_Server.Frames._2Body;
using Pr4_Server.Frames._3Footer;
using System.Text;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace Pr4_Server;

/// <summary>
/// Interaction logic for MainWindow.xaml
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