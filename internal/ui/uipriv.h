// 2 january 2016
#import <Cocoa/Cocoa.h>
#import <WebKit/WebKit.h>
#import <pthread.h>
#import "mach_override.h"
#import "ui.h"
#import "_cgo_export.h"

// autolayout.m
extern void addConstraint(NSView *view, NSString *constraint, NSDictionary *metrics, NSDictionary *views);
extern NSLayoutPriority horzHuggingPri(NSView *view);
extern NSLayoutPriority vertHuggingPri(NSView *view);
extern void setHuggingPri(NSView *view, NSLayoutPriority priority, NSLayoutConstraintOrientation orientation);
extern void layoutSingleView(NSView *superview, NSView *subview, int margined);
extern NSSize fittingAlignmentSize(NSView *view);
