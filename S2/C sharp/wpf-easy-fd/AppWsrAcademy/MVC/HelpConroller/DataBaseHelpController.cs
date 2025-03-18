using AppWsrAcademy.DataFilesApp;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AppWsrAcademy.MVC.HelpConroller
{
    class DataBaseHelpController
    {
        public static string GetLoginMain(string getSetLogin) {
            User userObj = OdbConnectHelper.entObj.User.FirstOrDefault(x => x.Login == getSetLogin);
            while (true) {
                if (userObj != null) {
                    return "Пользователь есть!";
                }
                else
                {
                    return "Такого пользователя нет!";
                }
            }
        }
    }
}
