#include <iostream>
#include <aws/core/Aws.h>
#include <aws/s3/S3Client.h>
#include <aws/s3/model/Bucket.h>

#ifdef _MSC_VER
#define FAL_EXPORT __declspec(dllexport)
#else
#define FAL_EXPORT __attribute__((visibility("default")))
#endif

int _listS3Buckets() {
    Aws::SDKOptions options;
    options.loggingOptions.logLevel = Aws::Utils::Logging::LogLevel::Debug;

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
    FAL_EXPORT int addAndMultiplies(int a) {
        return _addAndMultiples(a);
    }

    FAL_EXPORT int listS3Buckets() {
        return _listS3Buckets();
    }
}
