package paystack_test

import (
	"context"
	"log"
	"os"

	"github.com/chokey2nv/go-payment-gateways/paystack"
	"github.com/chokey2nv/go-payment-gateways/paystack/client"
	"github.com/chokey2nv/go-payment-gateways/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/joho/godotenv"
)

var _ = Describe("PlanService Integration (Paystack)", func() {
	_ = godotenv.Load(".env")
	var (
		// rec         *recorder.Recorder
		psClient    *client.PayStackClient
		planService *paystack.PlanService
		ctx         context.Context
		// planID      int
	)

	BeforeEach(func() {
		ctx = context.Background()

		// var err error
		// rec, err = recorder.New("fixtures/paystack_create_plan")
		// Expect(err).To(BeNil())
		// rec.SetMode(recorder.ModeReplayOnly)

		// httpClient := &http.Client{
		// 	Transport: rec,
		// }

		secretKey := os.Getenv("PAYSTACK_SECRET_KEY")

		log.Println(secretKey)

		Expect(secretKey).ToNot(BeEmpty())

		psClient = client.New(secretKey)
		// psClient = client.New(
		// 	secretKey,
		// 	client.WithHTTPClient(httpClient),
		// )

		planService = paystack.NewPlanService(psClient)
	})

	AfterEach(func() {
		// rec.Stop()
	})

	// It("should create plan on real paystack", func() {

	// 	res, err := planService.CreatePlan(ctx, models.CreatePlanRequest{
	// 		Name:     fmt.Sprintf("SDK Test %d", time.Now().Unix()),
	// 		Amount:   "50000",
	// 		Interval: "monthly",
	// 		Currency: "NGN",
	// 	})

	// 	Expect(err).To(BeNil())
	// 	Expect(res).NotTo(BeNil())
	// 	Expect(res.PlanCode).NotTo(BeEmpty())
	// 	if res != nil {
	// 		planID = res.ID
	// 	}
	// })
	// It("should fetch plan on real paystack", func() {
	// 	res, err := planService.FetchPlan(ctx, fmt.Sprintf("%d", planID))
	// 	Expect(err).To(BeNil())
	// 	Expect(res).NotTo(BeNil())
	// 	Expect(res.PlanCode).NotTo(BeEmpty())
	// })
	It("should list plan on real paystack", func() {
		res, meta, err := planService.ListPlan(ctx)
		utils.ErrorLog(res)
		utils.ErrorLog(meta)
		Expect(meta.Total).NotTo(BeNil())
		Expect(err).To(BeNil())
		Expect(res).NotTo(BeNil())
	})
})
