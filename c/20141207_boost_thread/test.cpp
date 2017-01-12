#include <boost/thread/thread.hpp>
#include <iostream>

class Hello {
public:
  int data;
  void hello(std::string name) {
    std::cout << "hello: " << name << ", " << data << std::endl;
  }
};


int main() {
  Hello* h = new Hello;
  h->data = 10;

  boost::thread th(boost::bind(&Hello::hello, h, "world"));
  th.join();

  
  return 0;
}
