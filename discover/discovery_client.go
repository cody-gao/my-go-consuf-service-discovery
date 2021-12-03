/**
 * @Author: gaoerpeng
 * @Description:
 * @Version: 1.0.0
 * @Date: 2021/11/30 11:41 上午
 */

package discover

import "log"

type DiscoveryClient interface {
	Register(serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) bool

	DeRegister(instanceId string, logger *log.Logger) bool

	DiscoverServices(serviceName string, logger *log.Logger) []interface{}
}
