using AppWsrAcademy.MVC.HelpConroller;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AppWsrAcademy.MVC.View
{
    class ViewLogin
    {
        public string GetLogin(string loginCheckOdb) {
            return DataBaseHelpController.GetLoginMain(loginCheckOdb);
        }
    }
}
