# Правильные комментарии


Продолжаем впитывать рекомендации

1. Необходимо писать максимально информативные комментарии, если они необходимы (регулярные выражения например)

Пример:

```
// Поиск по формату: kk:mm:ss EEE, MMM dd, yyyy 
Pattern timeMatcher = Pattern.compile( "\\d*:\\d*:\\d* \\w*, \\w* \\d*, \\d*");
```

2. В комментариях пишите о намерениях что-то сделать

```
// Мы пытаемся спровоцировать "состояние гонки", 
// создавая большое количество программных потоков.
```

3. Пишите предупреждения в комментариях, если они необходимы (например предупредить о том, что ту или иную функцию лучше не запускать)

```
// Не запускайте эти тесты, 
// потому что они будут работать долго (пару часов)
```

4. Оставляйте комментарии вида `TODO` для того функционала, который собираетесь дописать, но в будущем
5. Проясняйте загадочные аргументы или значения, которые импортируете из библиотек или используются в коде, который нельзя изменить, например, если они названы неясно, в остальных случаях все должно быть ясно в названии
6. Подчеркиваете важность того или иного функционала, если он очень важен

```
// Вызов trim() очень важен. Он удаляет начальные пробелы, 
// чтобы строка успешно интерпретировалась как список:
String listItemContent = match.group(3).trim();
```

Задания:

[7_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson15/7_example1.go)

```go
// ПУНКТ 1

// формирование отчета о зарплате за определенный период
func (employee *Employee) PrepareSalaryReportByPeriod(period time.Time) Report {}

// отношение между стадиями обработки и типами оборудований по торгам
type ProccesingStagesAndEquipmentTypes struct {
	// fields...
}

// отношение между оборудованием покупателей и поставщиков в тендере
type SellerAndVendorsEquipmentOfTender struct {
	// fields...
}

// история взаимодействия с компанией по тендеру
type HistoryInteractionWithCompanyOfTender struct {
	// fields...
}

// отношение описательной части рабочей части
// (процесс, когда тендер находится в обработке) тендера
type WorkPartDescriptionOfTender struct {
	// fields...
}

// ПУНКТ 2

// мы формируем коммерческое предложение по торгу
func GetCommericalProposal(object_id string, columns data_analyze.ColumnsValues, ndsStatus string) (Answer_from_kp, error) {
}

// мы ищем текстовые выражения по ключу в файле
func FindSentenceByKeyInFile(key string) string {
}

// ПУНКТ 3

type userManagerApi struct {
    // реализация сервера, который служит заглушкой
    // для еще не реализованных запросов
    user_manager_v1.UnimplementedUserManagerServer
    userManager UserManager
}

// ПУНКТ 4

// Не используйте этот метод при реализации обработчика запросов на сервере,
// так как он некорректно работает с базой данных
func AddTypeEquipment(ctx *fiber.Ctx) error {
}


// ПУНКТ 5

for _, row := range fileRows {
        // Данный вызов TrimSpace очень важен
        // здесь мы удаляем пробелы с концов строки,
        // чтобы далее она правильно интерпритировалась в массив
        row = strings.TrimSpace(row)
        pairKeyAndValue := strings.Split(row, "=")
        var keyOfPair, valueOfPair string = pairKeyAndValue[0], pairKeyAndValue[1]
        textStorage[keyOfPair] = valueOfPair
}

// ПУНКТ 6

func main() {
    MustSetUpEnv()
    log := SetupLoger("local")
    authConfig := auth_config.MustLoad()
    authApllication := auth_app.New(log, authConfig.GRPC.Port)
    authApllication.GRPCServer.MustRun()
    // TODO: все упаковать в gorutines и добавить канал,
    // который ожидает сигнала по завершению
    // затем GracefulShotDowmn
}

func (userManager *userManagerApi) GetUserProfile(
    ctx context.Context,
    request *user_manager_v1.GetUserProfileRequest,
) (*user_manager_v1.UserProfile, error) {
    // TODO: сделать валидацию данных
    user, errGetUser := userManager.userManager.GetUserProfile(ctx, request.GetUserId())

    if errGetUser != nil {
        return nil, errGetUser
    }
    
    return &user_manager_v1.UserProfile{
        Email:       user.Login,
        Name:        user.Name,
        Surname:     user.Surname,
        Sex:         user.Sex,
        Height:      user.Height,
        Weight:      user.Weight,
        DateBirthMs: user.BirthDateMs,
    }, nil
}

func (userManager *userManagerApi) GetMatchStatistic(
    context.Context,
    *user_manager_v1.GetUserStatisticRequest,
) (*user_manager_v1.UserStatistic, error) {
    // TODO: реализовать метод получения статистики по матчу
    return &user_manager_v1.UserStatistic{}, nil
}


func (userManager *userManagerApi) GetUserMatchesStatistic(
    context.Context,
    *user_manager_v1.GetUserStatisticRequest,
) (*user_manager_v1.UserMatchesStatistic, error) {
    // TODO: реализовать метод получения статистики по матчам
    return &user_manager_v1.UserMatchesStatistic{}, nil
}


func (userManager *userManagerApi) GetLeagueSchedule(
    context.Context,
    *user_manager_v1.GetLeagueScheduleRequest,
) (*user_manager_v1.LeagueSchedule, error) {
    // TODO: реализовать метод получения расписания лиги
    return &user_manager_v1.LeagueSchedule{}, nil
}


func (userManager *userManagerApi) GetLeagueStanding(
    context.Context,
    *user_manager_v1.GetLeagueStandingRequest,
) (*user_manager_v1.LeagueStanding, error) {
    // TODO: реализовать метод получения турнирной таблицы турнира
    return &user_manager_v1.LeagueStanding{}, nil

}
```

