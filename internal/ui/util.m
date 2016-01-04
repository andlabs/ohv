// 2 january 2016
#import "uipriv.h"

char *fromNSURL(id u)
{
	NSURL *url = (NSURL *) u;

	return strdup([[url absoluteString] UTF8String]);
}

id toNSURL(char *str)
{
	NSString *s;

	s = [NSString stringWithUTF8String:str];
	free(str);
	return [NSURL URLWithString:s];
}

id toFileNSURL(char *str, char *anchor)
{
	NSString *s;
	NSMutableURL *url;

	s = [NSString stringWithUTF8String:str];
	free(str);
	url = [NSMutableURL fileURLWithPath:s];
	if (anchor != NULL) {
		s = [NSString stringWithUTF8String:anchor];
		free(anchor);
		[url setFragment:s];
	}
	return url;
}
