using AppWsrAcademy.MVC.View;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AppWsrAcademy.MVC.Controller
{
    class ControllerLogin
    {
        public static string CheckLoginOdb(string loginCheck) {
            ViewLogin viewLogin = new ViewLogin();
            return viewLogin.GetLogin(loginCheck);
        }
    }
}
