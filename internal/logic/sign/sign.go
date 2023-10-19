package sign

import (
	"mpcServer/internal/logic/sign/util/li17"
	"mpcServer/internal/model"
	"sync"
)

func (a *sSigner) SignSendRequestP1(context1 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.SignSendRequestP1(context1)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) SignRecvRequestP2(context2 string, request string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.SignRecvRequestP2(context2, request)
		await.wg.Done()
	})
	return await
}
func (a *sSigner) SignSendPartialP2(context2, msg string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.SignSendPartialP2(context2, msg)
		await.wg.Done()
	})
	return await
}
func (a *sSigner) SignSendPartialP1(context1, sign2, msg string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.SignSendPartialP1(context1, sign2, msg)
		await.wg.Done()
	})
	return await
}
