// 2 january 2016
#import "uipriv.h"

id toNSURL(char *str)
{
	NSString *s;

	s = [NSString stringWithUTF8String:str];
	free(str);
	return [NSURL URLWithString:s];
}

id toFileNSURL(char *str)
{
	NSString *s;

	s = [NSString stringWithUTF8String:str];
	free(str);
	return [NSURL fileURLWithPath:s];
}
