#include <iostream>
// 出租方
class Lessor{
public:
       virtual void rentHouse()=0;
};
// 房东
class LandLord : public Lessor{
public:
       virtual void rentHouse(){
              std::cout << "房东出租房子" << std::endl;
       }
};
// 中介
class HouseProxy : public Lessor{
public:
       virtual void rentHouse(){
              std::cout << "中介找房东租房子" << std::endl;
              landlord.rentHouse();
       }
private:
       LandLord landlord;
};

//租客
class Renter{
public:
       void findHouse(){
              std::cout << "租客找中介租房子" << std::endl;
              lessor.rentHouse();
       }
private:
       HouseProxy lessor;
};
int main(){
       Renter renter;
       renter.findHouse();
       return 0;
}

