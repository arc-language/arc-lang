

// what this type
async func otherAsyncFunc2() int32 {

    // do some long running work or task
    sleep(10)

    return 1000
}

async func myAsyncFunction2() int32 {

    let someResult = await otherAsyncFunc2()
    printf("%d\n", someResult)

    return 0
}

func main() int32 {

    myAsyncFunction2()

    // would exit instantly cause it's async, so add sleep() or block it with something
    sleep(20)

    return 0
}



// compiler converts above into 

func otherAsyncFunc2() int32 {

    // do some long running work or task
    sleep(10)

    return 1000
}

func myAsyncFunction2() int32 {

    let someResult = get_task_result(otherAsyncFunc2, ... func args)
    printf("%d\n", someResult)

    return 0
}

func main() int32 {

    schedule_task(myAsyncFunction2, ... func args)

    // would exit instantly cause it's async, so add sleep() or block it with something
    sleep(20)

    return 0
}