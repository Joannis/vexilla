# Vexilla Client - Cocoa (iOS and OSX)

This is the Cocoa (iOS and OSX) client library for Vexilla, a static file based feature flag system.

## Getting Started

To get started is easy.

### Installation

Add VexillaClient as a dependency via Swift Packages
[https://developer.apple.com/documentation/xcode/adding_package_dependencies_to_your_app](https://developer.apple.com/documentation/xcode/adding_package_dependencies_to_your_app)



### Setup

You will need to create a Client within your app. This optionally takes in the `customInstanceHash` for use with gradual rollout.

After creation, call `fetchFlags`. This can be chained from the constructor since it returns the client instance.

```swift
var client = VexillaClient(environment: "dev", baseUrl: "https://BUCKET_NAME.s3-website-AWS_REGION.amazonaws.com", customInstanceHash: userId)

client.fetchFlags(fileName: "features.json", fetchCompletionHandler: { flags, error in

    guard flags != nil else {
        dump("Error: Flags were nil in test")
        return
    }

    client.setFlags(flags: flags!)
    
)}
```


### Generate Types (Optional)

We have created a tool to generate types for usage in your code. It will crawl your JSON structures and create consts or enums to help prevent typos and other "magic string" related issues. You just need to pass the URL of the JSON file where it is hosted.

#### NPM

To use the tool, you can run it directly from NPM.

```
npx vexilla types cocoa REMOTE_JSON_URL
```

#### Automatic Install Script

You can also use a precompiled binary for your platform. This install script will automatically choose the right binary for you:

```
curl -o- -s https://raw.githubusercontent.com/vexilla/vexilla/main/install.sh | bash
```

The install script also accepts a target install path:

```
curl -o- -s https://raw.githubusercontent.com/vexilla/vexilla/main/install.sh | bash
```

#### Manual Installation

If you prefer to download the binary manually you can get it from the releases section in Github, [https://github.com/vexilla/vexilla/releases](https://github.com/vexilla/vexilla/releases)

### Usage

Use the created client to check if a feature `Should` be on.

```swift
shouldGradual = client.should(FEATURE_NAME)
```


### Full Example

```swift
var client = VexillaClient(environment: "dev", baseUrl: "https://streamparrot-feature-flags.s3.amazonaws.com", customInstanceHash: "b7e91cc5-ec76-4ec3-9c1c-075032a13a1a")

client.fetchFlags(fileName: "features.json", fetchCompletionHandler: { flags, error in

    guard flags != nil else {
        dump("Error: Flags were nil in test")
        return
    }

    client.setFlags(flags: flags!)
    
    if client.should(FEATURE_NAME) {
        // Do the thing
    }
)}
```


## What are Feature Flags?

Feature flags are useful for enabling or disabling parts of your application without having to redeploy. In some cases, such as mobile applications, the redeploy could take up to a week.

See more about them here:

- [https://featureflags.io](https://featureflags.io)
- [https://en.wikipedia.org/wiki/Feature_toggle](https://en.wikipedia.org/wiki/Feature_toggle)

Feature Flags are also a fundamental building block for things such as A/B testing.

## How does it work?

The process is simple but has several steps to get up and running. Please see our in-depth guides in our [documentation](https://vexilla.dev/documentation).

## Contributing

Would you like to contribute to this client SDK? There are many ways you can help. Reporting issues or creating PRs are the most obvious. Helping triage the issues and PRs of others would also be a huge help. Being a vibrant member of the community on [Discord](https://discord.gg/GbJu3d93TC) is another way to help out.

If you would like to contribute to the app, docs, or other parts of the project, please go see our [Contribution Guide](https://vexilla.dev/documentation/contributing).

## Support

Have you run into a bug? Is there a feature you feel is missing? Feel free to create a [GitHub Issue](https://github.com/vexilla/vexilla/issues).

Another way to get support or help is to reach out in our [Discord community](https://discord.gg/GbJu3d93TC).

## Sponsors

No sponsors yet. This could be a link and icon for **your** company here.

## LICENSE

MIT

Copyright 2023 Vexilla

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.