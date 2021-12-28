#include <iostream>
#include <aws/core/Aws.h>
#include <aws/s3/S3Client.h>
#include <aws/s3/model/Bucket.h>

// __attribute__ is a GCC-specific keyword and does not exists on MS C++ Compiler
#ifdef _MSC_VER
#define __attribute__(x)
#endif

int _listS3Buckets() {
    Aws::SDKOptions options;
    Aws::InitAPI(options);
    {
        Aws::S3::S3Client client;

        auto outcome = client.ListBuckets();
        if (outcome.IsSuccess()) {
            std::cout << "Found " << outcome.GetResult().GetBuckets().size() << " buckets\n";
            for (auto&& b : outcome.GetResult().GetBuckets()) {
                std::cout << b.GetName() << std::endl;
            }
        }
        else {
            std::cout << "Failed with error: " << outcome.GetError() << std::endl;
        }
    }

    Aws::ShutdownAPI(options);
    return 0;
}

int _addAndMultiples(int num) {
    return (num + 1) * num;
}

extern "C" {
    int
    __attribute__((visibility("default")))
    addAndMultiplies(int a) {
        return _addAndMultiples(a);
    }

    int
    __attribute__((visibility("default")))
    listS3Buckets() {
        return _listS3Buckets();
    }
}
