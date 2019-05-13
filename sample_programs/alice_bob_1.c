#include <stdio.h>
#include <string.h>

int
main(int argc, char **argv)
{
	char alice_comment[16];
	char bob_comment[16];
	char *tmp;

	/*
	printf("&alice_comment[0]: %p\n", &alice_comment[0]);
	printf("&bob_comment[0]: %p\n", &bob_comment[0]);
	*/
	strcpy(alice_comment, "I hate Bob.");
	if (argc >= 2) {
		tmp = argv[1];
	} else {
		tmp = "I love Alice.";
	}
	strcpy(bob_comment, tmp);
	printf("Alice says \"%s\"\n", alice_comment);
	printf("Bob says \"%s\"\n", bob_comment);
}
