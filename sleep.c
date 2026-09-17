#include <IOKit/pwr_mgt/IOPMLib.h>
#include <CoreFoundation/CoreFoundation.h>

IOPMAssertionID createSleepAssertion() {
    IOPMAssertionID assertionID;

    CFStringRef reasonForActivity = CFStringCreateWithCString(
        kCFAllocatorDefault,
        "Go Anti-Sleep Tool Active",
        kCFStringEncodingUTF8
    );

    IOReturn success = IOPMAssertionCreateWithName(
        kIOPMAssertionTypePreventUserIdleDisplaySleep,
        kIOPMAssertionLevelOn,
        reasonForActivity,
        &assertionID
    );

    CFRelease(reasonForActivity);

    if (success == kIOReturnSuccess) {
        return assertionID;
    }
    return 0;
}

void releaseSleepAssertion(IOPMAssertionID assertionID) {
    if (assertionID != 0) {
        IOPMAssertionRelease(assertionID);
    }
}
