package sign

import (
	"mpcServer/internal/logic/sign/util/li17"
	"mpcServer/internal/model"
	"sync"
)

func (a *sSigner) KeygenSendHashProofP1(context1 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.KeygenSendHashProofP1(context1)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) KeygenRecvHashProofP2(context2, proof1 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.KeygenRecvHashProofP2(context2, proof1)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) KeygenSendZKProofP1(context1 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.KeygenSendZKProofP1(context1)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) KeygenRecvZKProofP1(context1, proof2 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.KeygenRecvZKProofP1(context1, proof2)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) KeygenSendZKProofP2(context1 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.KeygenSendZKProofP2(context1)
		await.wg.Done()
	})
	return await
}
func (a *sSigner) KeygenRecvZKProofP2(context2, proof1 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.KeygenRecvZKProofP2(context2, proof1)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) PublicKeyP1(context1 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.PublicKeyP1(context1)
		await.wg.Done()
	})
	return await
}
func (a *sSigner) PublicKeyP2(context2 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.PublicKeyP2(context2)
		await.wg.Done()
	})
	return await
}
