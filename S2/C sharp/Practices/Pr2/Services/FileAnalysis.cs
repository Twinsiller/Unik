using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Pr2.Services
{
    class FileAnalysis
    {
        string _filePath;
        string _fileName;
        int _countWords;
        int _countChars;

        public FileAnalysis(string filePath, int countWords, int countChars)
        {
            _filePath = filePath;
            _fileName = GetFileName(filePath);
            _countWords = countWords;
            _countChars = countChars;
        }

        string GetFileName(string filePath)
        {
            return Path.GetFileName(filePath);
        }
    }
}
