// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package opsworkscm provides the client and types for making API
// requests to AWS OpsWorks for Chef Automate.
//
// AWS OpsWorks for Chef Automate is a service that runs and manages configuration
// management servers.
//
// Glossary of terms
//
//    * Server: A configuration management server that can be highly-available.
//    The configuration manager runs on your instances by using various AWS
//    services, such as Amazon Elastic Compute Cloud (EC2), and potentially
//    Amazon Relational Database Service (RDS). A server is a generic abstraction
//    over the configuration manager that you want to use, much like Amazon
//    RDS. In AWS OpsWorks for Chef Automate, you do not start or stop servers.
//    After you create servers, they continue to run until they are deleted.
//
//    * Engine: The specific configuration manager that you want to use (such
//    as Chef) is the engine.
//
//    * Backup: This is an application-level backup of the data that the configuration
//    manager stores. A backup creates a .tar.gz file that is stored in an Amazon
//    Simple Storage Service (S3) bucket in your account. AWS OpsWorks for Chef
//    Automate creates the S3 bucket when you launch the first instance. A backup
//    maintains a snapshot of all of a server's important attributes at the
//    time of the backup.
//
//    * Events: Events are always related to a server. Events are written during
//    server creation, when health checks run, when backups are created, etc.
//    When you delete a server, the server's events are also deleted.
//
//    * AccountAttributes: Every account has attributes that are assigned in
//    the AWS OpsWorks for Chef Automate database. These attributes store information
//    about configuration limits (servers, backups, etc.) and your customer
//    account.
//
// Endpoints
//
// AWS OpsWorks for Chef Automate supports the following endpoints, all HTTPS.
// You must connect to one of the following endpoints. Chef servers can only
// be accessed or managed within the endpoint in which they are created.
//
//    * opsworks-cm.us-east-1.amazonaws.com
//
//    * opsworks-cm.us-west-2.amazonaws.com
//
//    * opsworks-cm.eu-west-1.amazonaws.com
//
// Throttling limits
//
// All API operations allow for five requests per second with a burst of 10
// requests per second.
//
// See https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01 for more information on this service.
//
// See opsworkscm package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/opsworkscm/
//
// Using the Client
//
// To use the client for AWS OpsWorks for Chef Automate you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := opsworkscm.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS OpsWorks for Chef Automate client OpsWorksCM for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/opsworkscm/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.AssociateNode(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("AssociateNode result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.AssociateNodeWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package opsworkscm
