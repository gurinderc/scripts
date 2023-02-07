int RoundToThousands(int vv) => (vv + (500 * (vv > 0 ? 1 : -1))) / 1000 * 1000;

Console.WriteLine(RoundToThousands(-747));//-1000
Console.WriteLine(RoundToThousands(-1501));//-2000

Console.WriteLine(RoundToThousands(747));//1000
Console.WriteLine(RoundToThousands(1501));//2000

Console.WriteLine(RoundToThousands(401));//0
Console.WriteLine(RoundToThousands(3401));//3000
