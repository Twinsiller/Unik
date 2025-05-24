using Pr3_Client.ClassOpt;
using Pr3_Client.Frames._1Header;
using Pr3_Client.Frames._2Body;
using Pr3_Client.Frames._3Footer;
using System.IO;
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

namespace Pr3_Client;

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


        // Например, запись на рабочий стол:
        string desktopPath = Environment.GetFolderPath(Environment.SpecialFolder.Desktop);
        string logFilePath = System.IO.Path.Combine(desktopPath, "log-client.txt");

        FileStream fs = new FileStream(logFilePath, FileMode.Create);
        StreamWriter sw = new StreamWriter(fs) { AutoFlush = true };
        Console.SetOut(sw);

        Console.WriteLine("Это сообщение записано в log-client.txt");
    }
}