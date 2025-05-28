using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Pr4_Server.ClassOpt
{
    public class GameLogic
    {
        public string SecretCode { get; private set; }

        public void GenerateCode()
        {
            var rand = new Random();
            SecretCode = new string(Enumerable.Range(0, 4)
                .Select(_ => (char)rand.Next('A', 'Z' + 1)).ToArray());
        }

        public (int black, int white) Evaluate(string guess)
        {
            int black = 0, white = 0;
            bool[] secretUsed = new bool[4];
            bool[] guessUsed = new bool[4];

            for (int i = 0; i < 4; i++)
            {
                if (guess[i] == SecretCode[i])
                {
                    black++;
                    secretUsed[i] = guessUsed[i] = true;
                }
            }

            for (int i = 0; i < 4; i++)
            {
                if (guessUsed[i]) continue;

                for (int j = 0; j < 4; j++)
                {
                    if (!secretUsed[j] && guess[i] == SecretCode[j])
                    {
                        white++;
                        secretUsed[j] = true;
                        break;
                    }
                }
            }

            return (black, white);
        }
    }


}
