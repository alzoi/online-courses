// На входе все числа, кроме одного, имеют пару. Найти число без пары.

#include <iostream>
#include <map>

int main() {

	// Словарь (таблица: ключ и значение для этого ключа).
	std::map<unsigned long long int, int> t_dat1;
	// Итератор по словарю.
	std::map<unsigned long long int, int>::iterator itr;
	unsigned long long int idata;
 
	// Получаем значения со стандартного ввода.
	while( std::cin.good() ){      
		
		std::cin >> idata;
		
		// Находим ключ в словаре.
		itr = t_dat1.find(idata);
		if(itr == t_dat1.end()) {
			// Добавляем элемент в словарь.
			t_dat1.insert(std::pair<unsigned long long int, int>(idata, 1));
		} else {
			// Если ключ есть увеличиваем значение на 1.
			t_dat1[idata]++;
		}
	}

	// Просматриваем записи в словаре.
	for (auto it = t_dat1.begin(); it != t_dat1.end(); it++) {
		
		// Если в словаре есть элемент со значением 1.
		if(it->second == 1) {
			std::cout << it->first;
			break;
		}

	}
}

/*

Ввод: 1 2 3 2 1 3 4 3
Вывод 4

*/
