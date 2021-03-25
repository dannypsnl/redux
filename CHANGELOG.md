## v2.2.2

remove the dependency about github.com/dannypsnl/assert since it already is removed

## v2.2.1

fix: support go module now

## v2.2.0

At before we use rematcher.Action(rematcher.ReducerMethod).With(payload) to generate our action for dispatching, now we just add a field with action tag, the format is: action:"ReducerMethodName"

## v2.1.2

fix:

* #52
* #53

improving performance by moving job to only executing once part.

## v2.1.1

improving store

## v2.1.0

feature: rematch, ref #48

## v2.0.3

fix: warning message for duplicating reducer typo

## v2.0.2

* remove readme warning of v2
* patch version to let code dependency can update

## v2.0.1

Fix bug #49, duplicated reducer will cause panic now

## v2.0.0

* The new store allows you:
  * using any type to be state & action.
  * let you can use the lambda that works as well as the global function by using their address instead of runtime name
* Follow first change, now you can send anything to Dispatch. It will pick matching reducer executing
* As we told, we using address rather than runtime name. The purpose of this change it let user can use function itself to get the mapping state! Now GetState be rename to StateOf. And it expected a reducer will be sent in.

<a name="v1.3.3"></a>
## [v1.3.3](https://github.com/dannypsnl/redux/compare/v1.3.2...v1.3.3)

> 2018-03-06

### Clean

* release v1.3.3
* update CHANGLOGS

### Fix

* Bug [#33](https://github.com/dannypsnl/redux/issues/33), repeat update the state when have middleware


<a name="v1.3.2"></a>
## [v1.3.2](https://github.com/dannypsnl/redux/compare/v1.3.1...v1.3.2)

> 2018-03-06

### Clean

* release v1.3.2 for bug fix
* udpate changelogs

### Fix

* expected should update state after call `next`


<a name="v1.3.1"></a>
## [v1.3.1](https://github.com/dannypsnl/redux/compare/v1.3.0...v1.3.1)

> 2018-03-06

### Clean

* release v1.3.1 for fix bug
* update changelogs
* format
* format
* update doc, add usage for ApplyMiddleware
* update changelogs
* move panic tests to independent file
* update `ApplyMiddleware` doc
* update readme
* Add package doc of middleware

### Fix

* ApplyMiddleware should return *Store


<a name="v1.3.0"></a>
## [v1.3.0](https://github.com/dannypsnl/redux/compare/v1.2.2...v1.3.0)

> 2018-03-06

### Add

* detect middleware API
* middleware impl

### Clean

* patch version for middleware, v1.3.0
* fix README typo, patch some text
* update Readme, and refactoring the code
* update readme
* generate changelogs
* remove left test code


<a name="v1.2.2"></a>
## [v1.2.2](https://github.com/dannypsnl/redux/compare/v1.2.1...v1.2.2)

> 2018-02-19

### Add

* patch ver 1.2.2

### Clean

* benchmark udpateState, and an reduced code ver
* improving comment
* abstract Dispatch's code
* test Dispatch in concurrency

### Fix

* README's error, `go get https://...`


<a name="v1.2.1"></a>
## [v1.2.1](https://github.com/dannypsnl/redux/compare/v1.2.0...v1.2.1)

> 2018-02-17

### Add

* [Release] ver 1.2.1
* more version in testing
* Tester for more work

### Clean

* remove benching unneed
* remove import path
* use native map implement, help it pass old version Go CI
* rename tests & refactor tests
* rename tests
* add more info into TestGetReducerName error message
* remove not useful script `run-example.sh`
* a better test name

### Delete

* 1.10 will be resolve as 1.1, tip -> master

### Test

* sync.Map & map's Loading speed


<a name="v1.2.0"></a>
## [v1.2.0](https://github.com/dannypsnl/redux/compare/v1.1.0...v1.2.0)

> 2018-01-22


<a name="v1.1.0"></a>
## [v1.1.0](https://github.com/dannypsnl/redux/compare/v1.0.3...v1.1.0)

> 2018-01-10


<a name="v1.0.3"></a>
## [v1.0.3](https://github.com/dannypsnl/redux/compare/v1.0.2...v1.0.3)

> 2018-01-06


<a name="v1.0.2"></a>
## [v1.0.2](https://github.com/dannypsnl/redux/compare/v1.0.1...v1.0.2)

> 2018-01-01


<a name="v1.0.1"></a>
## [v1.0.1](https://github.com/dannypsnl/redux/compare/v1.0.0...v1.0.1)

> 2017-12-31


<a name="v1.0.0"></a>
## [v1.0.0](https://github.com/dannypsnl/redux/compare/v0.7.1...v1.0.0)

> 2017-12-31


<a name="v0.7.1"></a>
## [v0.7.1](https://github.com/dannypsnl/redux/compare/v0.7.0...v0.7.1)

> 2017-12-31


<a name="v0.7.0"></a>
## [v0.7.0](https://github.com/dannypsnl/redux/compare/v0.6.0...v0.7.0)

> 2017-12-31


<a name="v0.6.0"></a>
## [v0.6.0](https://github.com/dannypsnl/redux/compare/v0.5.1...v0.6.0)

> 2017-12-30


<a name="v0.5.1"></a>
## [v0.5.1](https://github.com/dannypsnl/redux/compare/v0.5.0...v0.5.1)

> 2017-12-30


<a name="v0.5.0"></a>
## [v0.5.0](https://github.com/dannypsnl/redux/compare/v0.4.1...v0.5.0)

> 2017-12-30


<a name="v0.4.1"></a>
## [v0.4.1](https://github.com/dannypsnl/redux/compare/v0.4.0...v0.4.1)

> 2017-12-30


<a name="v0.4.0"></a>
## [v0.4.0](https://github.com/dannypsnl/redux/compare/v0.3.3...v0.4.0)

> 2017-12-27


<a name="v0.3.3"></a>
## [v0.3.3](https://github.com/dannypsnl/redux/compare/v0.3.2...v0.3.3)

> 2017-12-24


<a name="v0.3.2"></a>
## [v0.3.2](https://github.com/dannypsnl/redux/compare/v0.3.1...v0.3.2)

> 2017-12-23


<a name="v0.3.1"></a>
## [v0.3.1](https://github.com/dannypsnl/redux/compare/v0.3.0...v0.3.1)

> 2017-12-23


<a name="v0.3.0"></a>
## [v0.3.0](https://github.com/dannypsnl/redux/compare/v0.2.5...v0.3.0)

> 2017-12-19


<a name="v0.2.5"></a>
## [v0.2.5](https://github.com/dannypsnl/redux/compare/v0.2.4...v0.2.5)

> 2017-12-16


<a name="v0.2.4"></a>
## [v0.2.4](https://github.com/dannypsnl/redux/compare/v0.2.3...v0.2.4)

> 2017-12-11


<a name="v0.2.3"></a>
## [v0.2.3](https://github.com/dannypsnl/redux/compare/v0.2.2...v0.2.3)

> 2017-12-11


<a name="v0.2.2"></a>
## [v0.2.2](https://github.com/dannypsnl/redux/compare/v0.2.1...v0.2.2)

> 2017-12-05


<a name="v0.2.1"></a>
## [v0.2.1](https://github.com/dannypsnl/redux/compare/v0.2.0...v0.2.1)

> 2017-12-05


<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/dannypsnl/redux/compare/v0.1.7...v0.2.0)

> 2017-12-04


<a name="v0.1.7"></a>
## [v0.1.7](https://github.com/dannypsnl/redux/compare/v0.1.6...v0.1.7)

> 2017-12-02


<a name="v0.1.6"></a>
## [v0.1.6](https://github.com/dannypsnl/redux/compare/v0.1.5...v0.1.6)

> 2017-12-02


<a name="v0.1.5"></a>
## [v0.1.5](https://github.com/dannypsnl/redux/compare/v0.1.4...v0.1.5)

> 2017-12-02


<a name="v0.1.4"></a>
## [v0.1.4](https://github.com/dannypsnl/redux/compare/v0.1.3...v0.1.4)

> 2017-12-02


<a name="v0.1.3"></a>
## [v0.1.3](https://github.com/dannypsnl/redux/compare/v0.1.2...v0.1.3)

> 2017-12-02


<a name="v0.1.2"></a>
## [v0.1.2](https://github.com/dannypsnl/redux/compare/v0.1.1...v0.1.2)

> 2017-12-02


<a name="v0.1.1"></a>
## [v0.1.1](https://github.com/dannypsnl/redux/compare/v0.1.0...v0.1.1)

> 2017-12-02


<a name="v0.1.0"></a>
## v0.1.0

> 2017-12-02

