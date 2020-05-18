// Длинная арифметика - сложение длинных чисел.

#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

int main() {

	std::string num1, num2, str;
	std::vector<int> vnum1, vnum2;
	int carr, i, base, BASE = 1000000000;

	// Получаем числа.
	std::cin >> num1 >> num2;
	
	base = 9;
	// Считываем числа в массив.
	str = num1;
	for (i =(int)str.length(); i > 0; i -= base) {
		if (i < base){
			vnum1.push_back(atoi(str.substr(0, i).c_str()));
		} else {
			vnum1.push_back(atoi(str.substr(i-base, base).c_str()));
		}
	}
	str = num2;
	for (i =(int)str.length(); i > 0; i -= base) {
		if (i < base){
			vnum2.push_back(atoi(str.substr(0, i).c_str()));
		} else {
			vnum2.push_back(atoi(str.substr(i-base, base).c_str()));
		}
	}

	// Сложение vnum1 = vnum1 + vnum2.
	carr = 0;
	for (i = 0; i < std::max(vnum1.size(), vnum2.size()) || carr !=0; ++i) {
		if (i == vnum1.size()) {
			vnum1.push_back(0);
		}
		vnum1[i] += carr + (i < vnum2.size() ? vnum2[i] : 0);
		carr = vnum1[i] >= BASE;
		
		if (carr != 0) {
			vnum1[i] -= BASE;
		}
	}

	// Результат.
	printf("%d", vnum1.empty() ? 0 : vnum1.back());
	for (i = (int)vnum1.size() - 2; i >= 0; --i) {
		printf("%09d", vnum1[i]);
	}
}
