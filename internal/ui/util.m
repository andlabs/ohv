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

// thanks to mikeash in irc.freenode.net/#macdev
// TODO requires 10.9; if we export this it needs to change
id toFileNSURL(char *str, char *anchor)
{
	NSString *s;
	NSURL *url;
	NSURLComponents *c;

	s = [NSString stringWithUTF8String:str];
	free(str);
	url = [NSURL fileURLWithPath:s];
	// TODO pass YES instead?
	c = [NSURLComponents componentsWithURL:url resolvingAgainstBaseURL:NO];
	if (anchor != NULL) {
		s = [NSString stringWithUTF8String:anchor];
		free(anchor);
		[c setFragment:s];
	}
	return [c URL];
}
