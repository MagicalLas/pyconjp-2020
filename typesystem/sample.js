const g = f => function* () {
    yield f(yield)(yield)
}()

const w = g => f => {
    const gen = g(f)
    gen.next()
    return gen
}

const n = g => a => b => {
    g.next(a)
    return g.next(b)
}

const add = a => b => a + b
const mul = a => b => a * b

const w2 = f => a => b => n(w(g)(f))(a)(b).value

console.log(w2(add)(1)(3))
console.log(w2(mul)(1)(3))

console.log(w2(add)(1))