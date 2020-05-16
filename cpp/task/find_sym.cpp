// Поиск в последовательности суммы двух чисел, которая равна заданному числу.
// 8
// 1 3 4 5
// Вывод 1.

#include <vector>
#include <fstream>
#include <sstream>
#include <iostream>
#include <set>

int main() {

    std::ifstream      file("input.txt");
    std::ofstream      file_out;
    std::string        line;
    int                target, value, sum;
    std::multiset<int> mst;

    file_out.open("output.txt");

    std::getline(file, line);
    target = std::stoi(line);
    
    std::getline(file, line);
    file.close();
    std::stringstream lineStream(line);

    while (lineStream >> value) {
        if (value < target) {
            mst.insert(value);
        }        
    }

    if(!mst.size()) {
        file_out << "0";
        file_out.close();
        return 0;
    }

    auto it1 = mst.begin();
    auto it2 = mst.end();
    it2--;

    while (it1 != it2) {
        sum = *it1 + *it2;
        if (sum < target) {
            it1++;
        } else if (sum > target) {
            it2--;
        } else {
            file_out << "1";
            file_out.close();
            return 0;
        }
    }

    file_out << "0";
    file_out.close();
    return 0;
}
