package mongo_test

import (
	"github.com/golang/mock/gomock"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"

	. "github.com/topfreegames/go-extensions-mongo"
	"github.com/topfreegames/go-extensions-mongo/interfaces"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mongo", func() {
	var config *viper.Viper
	var mockCtrl *gomock.Controller
	var mockDb *interfaces.MockMongoDB
	var mockColl *interfaces.MockCollection

	BeforeEach(func() {
		config = viper.New()
		config.SetConfigFile("../config/test.yaml")
		Expect(config.ReadInConfig()).NotTo(HaveOccurred())
	})

	Describe("[Unit]", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockDb = interfaces.NewMockMongoDB(mockCtrl)
			mockColl = interfaces.NewMockCollection(mockCtrl)
		})

		AfterEach(func() {
			mockCtrl.Finish()
		})

		Describe("Connect", func() {
			It("Should use config to load connection details", func() {
				_, err := NewClient("extensions.mongo", config, mockDb)
				Expect(err).NotTo(HaveOccurred())
			})

			It("Should call dial with timeout if timeout is specified", func() {
				config.Set("extensions.mongo.url", "localhost:80")
				client, err := NewClient("extensions.mongo", config)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("no reachable servers"))
				Expect(client).To(BeNil())
			})
		})

		Describe("Close", func() {
			It("Should close after creating", func() {
				client, err := NewClient("extensions.mongo", config, mockDb)
				Expect(err).NotTo(HaveOccurred())

				mockDb.EXPECT().Close()
				client.Close()
			})
		})

		Describe("Run", func() {
			It("Should execute command with run", func() {
				collectionName := "coll"

				client, err := NewClient("extensions.mongo", config, mockDb)
				Expect(err).NotTo(HaveOccurred())

				mockDb.EXPECT().Close()
				mockDb.EXPECT().C(collectionName).Return(mockColl, nil)
				mockColl.EXPECT().Insert(gomock.Any())

				c, _ := client.MongoDB.C(collectionName)
				err = c.Insert(bson.M{})
				Expect(err).NotTo(HaveOccurred())
				client.Close()
			})

			It("Should execute Run command", func() {
				client, err := NewClient("extensions.mongo", config, mockDb)
				Expect(err).NotTo(HaveOccurred())

				mockDb.EXPECT().Run(gomock.Any(), gomock.Any())

				var result string
				err = client.MongoDB.Run(bson.D{
					{"create", "mycollection"},
				}, &result)
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})