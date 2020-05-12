// A year on Jupiter ( the time it takes Jupiter to make one full rotation around the sun) takes about 12 Earth years. Write a program
// that converts Jovian years to earth years. Have the user specify the number of Jovian years. Allow fractional years.

#include <std.io.h>

int main (void)
{
  	double jovian, earth;
  	char leavingNow;

  	leavingNow = ''

  	while leavingNow != 'exit';
		printf("Please enter a number greater than 0, this will represent Jovian years\n");
		printf("Then I can tell you how many Earth years that will be \n");
		scanf(%lf, &jovian);
		  
		earth = jovian*12;
		printf("You're looking at %s Earth years. Wow!\n", earth);
		printf("If you would like to stop, type 'exit'");
		scanf(%c, &leavingNow);
}
