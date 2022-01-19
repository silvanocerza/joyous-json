# Test instructions

In order to give some context, we’re going to put ourselves in the following imaginary situation: you’re part of the DevX team and you’ve been asked to build a solution based on the following specs.

To deal with huge amounts of data, we need to build a swiss army knife to cut through a large stream of JSON objects. This tool is going to be used by other developers, data engineers and so on.

- _In broad terms it looks like in the attached image "Json Example", please have a look._

To give some background, the stream of JSON objects can be event logs from various operations on repository, infrastructure, forming an audit trail. To illustrate, let's take this imaginary scenario:

_Team X wants to update yesterday's dump of events, making it so only the events that concern their team have a specific key renamed and also insert a specific integer ID corresponding to the incident that they went through. Finally, they'll take the result and send it to the Sec team._

In concrete terms, this means that at least the following operations must be directly supported:

- Rejecting an object based on the value of a specific field.
- Retaining an object only if it has a field with a specific value.
- Adding a key value pair on an object
- Prefixing a key with a string

That stream is to be consumed from:

- the command line
  - where the stream comes from STDIN and is outputted on STDOUT
  - ex: cat locations.json_dump | <your tool>
  - feel free to design its usage to your liking
- In a second time, as a library
  - where we can configure the input source and output destination, such as reading from the network and outputting in a file for example

# Prerequisites

- Provide a script showcasing the scenario from Team X: running the script should give the expected output.
  - Their team is called "team-x"
  - The incident field to insert on each object is named "incident_id"
  - The value of their specific incident is 6502.
- handle all failure cases
- your solution should be simple to use

- **don’t bother with nested fields**

- it should be simple for library users to add new operations
- documentation
  - it should be easy for a user to understand how to use the cli tool
  - it should be easy for a user of the library to understand how to use its public api
  - provide a proper README.md:
    - explain how to setup and run your code
    - including all informations you feel that may be useful for a seamless coworker on-boarding
- tests

## Workflow

- Use Go to write your solution.
- Create a local git repository (we strongly advise you to make multiple commits, especially if you want to highlight your iterative process).
- Send us back the zip once you've completed the exercise or ran out of time.

## Optional objectives

These are to be completed if and only you've completed the mandatory requirements listed above. They won't be considered otherwise.

- Provide a way to compare values (less or greater than, longer or shorter than, …) rather than just equality or difference.
- Provide a way to apply operations only on objects only if they match a given predicate.
