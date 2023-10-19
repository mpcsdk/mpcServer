package sign

import (
	"mpcServer/internal/logic/sign/util/li17"
	"mpcServer/internal/model"
	"sync"
)

func (a *sSigner) GenContextP1(preivateKey, publicKey string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.GenContextP1(preivateKey, publicKey)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) GenContextP2(preivateKey, publicKey string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.GenContextP2(preivateKey, publicKey)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) SendZKProofP1(p1 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.SendZKProofP1(p1)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) RecvZKProofP1(p1, ZKProof2 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.RecvZKProofP1(p1, ZKProof2)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) SendZKProofP2(p2 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.SendZKProofP2(p2)
		await.wg.Done()
	})
	return await
}

func (a *sSigner) RecvZKProofP2(p2, ZKProof1 string) model.ISignerPromise {
	await := &signeraWait{
		wg: sync.WaitGroup{},
	}
	await.wg.Add(1)
	a.AddTask(func() {
		await.str = li17.RecvZKProofP2(p2, ZKProof1)
		await.wg.Done()
	})
	return await
}
