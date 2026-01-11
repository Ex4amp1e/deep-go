package homework6

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		copy(person.name[:], name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.params = paramsLayout.mana.set(person.params, mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.params = paramsLayout.health.set(person.params, health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.params = paramsLayout.respect.set(person.params, respect)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.params = paramsLayout.strength.set(person.params, strength)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.expLevel = expLevelLayout.experience.set(person.expLevel, experience)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.expLevel = expLevelLayout.level.set(person.expLevel, level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.setFlag(house)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.setFlag(house)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.setFlag(family)
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.params = paramsLayout.personType.set(person.params, personType)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

const (
	usernameLength = 42

	// Flags
	house  = 1
	gun    = 2
	family = 4
)

// 64 byte
type GamePerson struct {
	x        int32                // 4 byte
	y        int32                // 4 byte
	z        int32                // 4 byte
	gold     uint32               // 4 byte
	params   uint32               // 4 byte
	flags    uint8                // 1 byte
	expLevel uint8                // 1 byte
	name     [usernameLength]byte // 1*42 byte
}

type bitField32 = bitField[uint32]
type bitField8 = bitField[uint8]

var paramsLayout = struct {
	mana       bitField32
	health     bitField32
	respect    bitField32
	strength   bitField32
	personType bitField32
}{
	mana:       newBitField32(0, 10),
	health:     newBitField32(10, 10),
	respect:    newBitField32(20, 4),
	strength:   newBitField32(24, 4),
	personType: newBitField32(28, 4),
}

var expLevelLayout = struct {
	experience bitField8
	level      bitField8
}{
	experience: newBitField8(0, 4),
	level:      newBitField8(4, 4),
}

type bitField[T uint32 | uint8] struct {
	offset uint8
	size   uint8
	mask   T
}

func newBitField32(offset, size uint8) bitField32 {
	return newBitField[uint32](offset, size)
}

func newBitField8(offset, size uint8) bitField8 {
	return newBitField[uint8](offset, size)
}

func newBitField[T uint8 | uint32](offset, size uint8) bitField[T] {
	return bitField[T]{
		offset: offset,
		size:   size,
		mask:   (1 << size) - 1,
	}
}

func (bf *bitField[T]) get(source T) int {
	return int((source >> bf.offset) & bf.mask)
}

func (bf *bitField[T]) set(source T, value int) T {
	return (source & ^(bf.mask << bf.offset)) | (T(value)&bf.mask)<<bf.offset
}

func (p *GamePerson) hasFlag(flag uint8) bool {
	return p.flags&flag != 0
}

func (p *GamePerson) setFlag(flag uint8) {
	p.flags |= flag
}

func NewGamePerson(options ...Option) GamePerson {
	var gamePerson GamePerson
	for _, option := range options {
		option(&gamePerson)
	}
	return gamePerson
}

func (p *GamePerson) Name() string {
	return string(p.name[:])
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return paramsLayout.mana.get(p.params)
}

func (p *GamePerson) Health() int {
	return paramsLayout.health.get(p.params)
}

func (p *GamePerson) Respect() int {
	return paramsLayout.respect.get(p.params)
}

func (p *GamePerson) Strength() int {
	return paramsLayout.strength.get(p.params)
}

func (p *GamePerson) Experience() int {
	return expLevelLayout.experience.get(p.expLevel)
}

func (p *GamePerson) Level() int {
	return expLevelLayout.level.get(p.expLevel)
}

func (p *GamePerson) HasHouse() bool {
	return p.hasFlag(house)
}

func (p *GamePerson) HasGun() bool {
	return p.hasFlag(gun)
}

func (p *GamePerson) HasFamily() bool {
	return p.hasFlag(family)
}

func (p *GamePerson) Type() int {
	return paramsLayout.personType.get(p.params)
}
