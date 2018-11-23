// Automatically generated by MockGen. DO NOT EDIT!
// Source: db.go

package db

import (
	gomock "github.com/rafrombrc/gomock/gomock"
	gorm "github.com/jinzhu/gorm"
	dao "github.com/goodrain/rainbond/db/dao"
)

// Mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *_MockManagerRecorder
}

// Recorder for MockManager (not exported)
type _MockManagerRecorder struct {
	mock *MockManager
}

func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &_MockManagerRecorder{mock}
	return mock
}

func (_m *MockManager) EXPECT() *_MockManagerRecorder {
	return _m.recorder
}

func (_m *MockManager) CloseManager() error {
	ret := _m.ctrl.Call(_m, "CloseManager")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) CloseManager() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloseManager")
}

func (_m *MockManager) Begin() *gorm.DB {
	ret := _m.ctrl.Call(_m, "Begin")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (_mr *_MockManagerRecorder) Begin() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Begin")
}

func (_m *MockManager) LicenseDao() dao.LicenseDao {
	ret := _m.ctrl.Call(_m, "LicenseDao")
	ret0, _ := ret[0].(dao.LicenseDao)
	return ret0
}

func (_mr *_MockManagerRecorder) LicenseDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LicenseDao")
}

func (_m *MockManager) AppDao() dao.AppDao {
	ret := _m.ctrl.Call(_m, "AppDao")
	ret0, _ := ret[0].(dao.AppDao)
	return ret0
}

func (_mr *_MockManagerRecorder) AppDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AppDao")
}

func (_m *MockManager) TenantDao() dao.TenantDao {
	ret := _m.ctrl.Call(_m, "TenantDao")
	ret0, _ := ret[0].(dao.TenantDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantDao")
}

func (_m *MockManager) TenantDaoTransactions(db *gorm.DB) dao.TenantDao {
	ret := _m.ctrl.Call(_m, "TenantDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantDaoTransactions", arg0)
}

func (_m *MockManager) TenantServiceDao() dao.TenantServiceDao {
	ret := _m.ctrl.Call(_m, "TenantServiceDao")
	ret0, _ := ret[0].(dao.TenantServiceDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceDao")
}

func (_m *MockManager) TenantServiceDeleteDao() dao.TenantServiceDeleteDao {
	ret := _m.ctrl.Call(_m, "TenantServiceDeleteDao")
	ret0, _ := ret[0].(dao.TenantServiceDeleteDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceDeleteDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceDeleteDao")
}

func (_m *MockManager) TenantServiceDaoTransactions(db *gorm.DB) dao.TenantServiceDao {
	ret := _m.ctrl.Call(_m, "TenantServiceDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantServiceDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceDaoTransactions", arg0)
}

func (_m *MockManager) TenantServiceDeleteDaoTransactions(db *gorm.DB) dao.TenantServiceDeleteDao {
	ret := _m.ctrl.Call(_m, "TenantServiceDeleteDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantServiceDeleteDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceDeleteDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceDeleteDaoTransactions", arg0)
}

func (_m *MockManager) TenantServicesPortDao() dao.TenantServicesPortDao {
	ret := _m.ctrl.Call(_m, "TenantServicesPortDao")
	ret0, _ := ret[0].(dao.TenantServicesPortDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServicesPortDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServicesPortDao")
}

func (_m *MockManager) TenantServicesPortDaoTransactions(_param0 *gorm.DB) dao.TenantServicesPortDao {
	ret := _m.ctrl.Call(_m, "TenantServicesPortDaoTransactions", _param0)
	ret0, _ := ret[0].(dao.TenantServicesPortDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServicesPortDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServicesPortDaoTransactions", arg0)
}

func (_m *MockManager) TenantServiceRelationDao() dao.TenantServiceRelationDao {
	ret := _m.ctrl.Call(_m, "TenantServiceRelationDao")
	ret0, _ := ret[0].(dao.TenantServiceRelationDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceRelationDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceRelationDao")
}

func (_m *MockManager) TenantServiceRelationDaoTransactions(_param0 *gorm.DB) dao.TenantServiceRelationDao {
	ret := _m.ctrl.Call(_m, "TenantServiceRelationDaoTransactions", _param0)
	ret0, _ := ret[0].(dao.TenantServiceRelationDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceRelationDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceRelationDaoTransactions", arg0)
}

func (_m *MockManager) TenantServiceEnvVarDao() dao.TenantServiceEnvVarDao {
	ret := _m.ctrl.Call(_m, "TenantServiceEnvVarDao")
	ret0, _ := ret[0].(dao.TenantServiceEnvVarDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceEnvVarDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceEnvVarDao")
}

func (_m *MockManager) TenantServiceEnvVarDaoTransactions(_param0 *gorm.DB) dao.TenantServiceEnvVarDao {
	ret := _m.ctrl.Call(_m, "TenantServiceEnvVarDaoTransactions", _param0)
	ret0, _ := ret[0].(dao.TenantServiceEnvVarDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceEnvVarDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceEnvVarDaoTransactions", arg0)
}

func (_m *MockManager) TenantServiceMountRelationDao() dao.TenantServiceMountRelationDao {
	ret := _m.ctrl.Call(_m, "TenantServiceMountRelationDao")
	ret0, _ := ret[0].(dao.TenantServiceMountRelationDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceMountRelationDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceMountRelationDao")
}

func (_m *MockManager) TenantServiceMountRelationDaoTransactions(db *gorm.DB) dao.TenantServiceMountRelationDao {
	ret := _m.ctrl.Call(_m, "TenantServiceMountRelationDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantServiceMountRelationDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceMountRelationDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceMountRelationDaoTransactions", arg0)
}

func (_m *MockManager) TenantServiceVolumeDao() dao.TenantServiceVolumeDao {
	ret := _m.ctrl.Call(_m, "TenantServiceVolumeDao")
	ret0, _ := ret[0].(dao.TenantServiceVolumeDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceVolumeDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceVolumeDao")
}

func (_m *MockManager) TenantServiceVolumeDaoTransactions(_param0 *gorm.DB) dao.TenantServiceVolumeDao {
	ret := _m.ctrl.Call(_m, "TenantServiceVolumeDaoTransactions", _param0)
	ret0, _ := ret[0].(dao.TenantServiceVolumeDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceVolumeDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceVolumeDaoTransactions", arg0)
}

func (_m *MockManager) ServiceProbeDao() dao.ServiceProbeDao {
	ret := _m.ctrl.Call(_m, "ServiceProbeDao")
	ret0, _ := ret[0].(dao.ServiceProbeDao)
	return ret0
}

func (_mr *_MockManagerRecorder) ServiceProbeDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceProbeDao")
}

func (_m *MockManager) ServiceProbeDaoTransactions(_param0 *gorm.DB) dao.ServiceProbeDao {
	ret := _m.ctrl.Call(_m, "ServiceProbeDaoTransactions", _param0)
	ret0, _ := ret[0].(dao.ServiceProbeDao)
	return ret0
}

func (_mr *_MockManagerRecorder) ServiceProbeDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceProbeDaoTransactions", arg0)
}

func (_m *MockManager) TenantServiceLBMappingPortDao() dao.TenantServiceLBMappingPortDao {
	ret := _m.ctrl.Call(_m, "TenantServiceLBMappingPortDao")
	ret0, _ := ret[0].(dao.TenantServiceLBMappingPortDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceLBMappingPortDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceLBMappingPortDao")
}

func (_m *MockManager) TenantServiceLBMappingPortDaoTransactions(_param0 *gorm.DB) dao.TenantServiceLBMappingPortDao {
	ret := _m.ctrl.Call(_m, "TenantServiceLBMappingPortDaoTransactions", _param0)
	ret0, _ := ret[0].(dao.TenantServiceLBMappingPortDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceLBMappingPortDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceLBMappingPortDaoTransactions", arg0)
}

func (_m *MockManager) TenantServiceLabelDao() dao.TenantServiceLabelDao {
	ret := _m.ctrl.Call(_m, "TenantServiceLabelDao")
	ret0, _ := ret[0].(dao.TenantServiceLabelDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceLabelDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceLabelDao")
}

func (_m *MockManager) TenantServiceLabelDaoTransactions(db *gorm.DB) dao.TenantServiceLabelDao {
	ret := _m.ctrl.Call(_m, "TenantServiceLabelDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantServiceLabelDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServiceLabelDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServiceLabelDaoTransactions", arg0)
}

func (_m *MockManager) LocalSchedulerDao() dao.LocalSchedulerDao {
	ret := _m.ctrl.Call(_m, "LocalSchedulerDao")
	ret0, _ := ret[0].(dao.LocalSchedulerDao)
	return ret0
}

func (_mr *_MockManagerRecorder) LocalSchedulerDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LocalSchedulerDao")
}

func (_m *MockManager) TenantPluginDaoTransactions(db *gorm.DB) dao.TenantPluginDao {
	ret := _m.ctrl.Call(_m, "TenantPluginDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantPluginDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantPluginDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantPluginDaoTransactions", arg0)
}

func (_m *MockManager) TenantPluginDao() dao.TenantPluginDao {
	ret := _m.ctrl.Call(_m, "TenantPluginDao")
	ret0, _ := ret[0].(dao.TenantPluginDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantPluginDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantPluginDao")
}

func (_m *MockManager) TenantPluginDefaultENVDaoTransactions(db *gorm.DB) dao.TenantPluginDefaultENVDao {
	ret := _m.ctrl.Call(_m, "TenantPluginDefaultENVDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantPluginDefaultENVDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantPluginDefaultENVDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantPluginDefaultENVDaoTransactions", arg0)
}

func (_m *MockManager) TenantPluginDefaultENVDao() dao.TenantPluginDefaultENVDao {
	ret := _m.ctrl.Call(_m, "TenantPluginDefaultENVDao")
	ret0, _ := ret[0].(dao.TenantPluginDefaultENVDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantPluginDefaultENVDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantPluginDefaultENVDao")
}

func (_m *MockManager) TenantPluginBuildVersionDao() dao.TenantPluginBuildVersionDao {
	ret := _m.ctrl.Call(_m, "TenantPluginBuildVersionDao")
	ret0, _ := ret[0].(dao.TenantPluginBuildVersionDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantPluginBuildVersionDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantPluginBuildVersionDao")
}

func (_m *MockManager) TenantPluginBuildVersionDaoTransactions(db *gorm.DB) dao.TenantPluginBuildVersionDao {
	ret := _m.ctrl.Call(_m, "TenantPluginBuildVersionDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantPluginBuildVersionDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantPluginBuildVersionDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantPluginBuildVersionDaoTransactions", arg0)
}

func (_m *MockManager) TenantPluginVersionENVDao() dao.TenantPluginVersionEnvDao {
	ret := _m.ctrl.Call(_m, "TenantPluginVersionENVDao")
	ret0, _ := ret[0].(dao.TenantPluginVersionEnvDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantPluginVersionENVDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantPluginVersionENVDao")
}

func (_m *MockManager) TenantPluginVersionENVDaoTransactions(db *gorm.DB) dao.TenantPluginVersionEnvDao {
	ret := _m.ctrl.Call(_m, "TenantPluginVersionENVDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantPluginVersionEnvDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantPluginVersionENVDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantPluginVersionENVDaoTransactions", arg0)
}

func (_m *MockManager) TenantServicePluginRelationDao() dao.TenantServicePluginRelationDao {
	ret := _m.ctrl.Call(_m, "TenantServicePluginRelationDao")
	ret0, _ := ret[0].(dao.TenantServicePluginRelationDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServicePluginRelationDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServicePluginRelationDao")
}

func (_m *MockManager) TenantServicePluginRelationDaoTransactions(db *gorm.DB) dao.TenantServicePluginRelationDao {
	ret := _m.ctrl.Call(_m, "TenantServicePluginRelationDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantServicePluginRelationDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServicePluginRelationDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServicePluginRelationDaoTransactions", arg0)
}

func (_m *MockManager) TenantServicesStreamPluginPortDao() dao.TenantServicesStreamPluginPortDao {
	ret := _m.ctrl.Call(_m, "TenantServicesStreamPluginPortDao")
	ret0, _ := ret[0].(dao.TenantServicesStreamPluginPortDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServicesStreamPluginPortDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServicesStreamPluginPortDao")
}

func (_m *MockManager) TenantServicesStreamPluginPortDaoTransactions(db *gorm.DB) dao.TenantServicesStreamPluginPortDao {
	ret := _m.ctrl.Call(_m, "TenantServicesStreamPluginPortDaoTransactions", db)
	ret0, _ := ret[0].(dao.TenantServicesStreamPluginPortDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TenantServicesStreamPluginPortDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TenantServicesStreamPluginPortDaoTransactions", arg0)
}

func (_m *MockManager) CodeCheckResultDao() dao.CodeCheckResultDao {
	ret := _m.ctrl.Call(_m, "CodeCheckResultDao")
	ret0, _ := ret[0].(dao.CodeCheckResultDao)
	return ret0
}

func (_mr *_MockManagerRecorder) CodeCheckResultDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CodeCheckResultDao")
}

func (_m *MockManager) CodeCheckResultDaoTransactions(db *gorm.DB) dao.CodeCheckResultDao {
	ret := _m.ctrl.Call(_m, "CodeCheckResultDaoTransactions", db)
	ret0, _ := ret[0].(dao.CodeCheckResultDao)
	return ret0
}

func (_mr *_MockManagerRecorder) CodeCheckResultDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CodeCheckResultDaoTransactions", arg0)
}

func (_m *MockManager) ServiceEventDao() dao.EventDao {
	ret := _m.ctrl.Call(_m, "ServiceEventDao")
	ret0, _ := ret[0].(dao.EventDao)
	return ret0
}

func (_mr *_MockManagerRecorder) ServiceEventDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceEventDao")
}

func (_m *MockManager) ServiceEventDaoTransactions(db *gorm.DB) dao.EventDao {
	ret := _m.ctrl.Call(_m, "ServiceEventDaoTransactions", db)
	ret0, _ := ret[0].(dao.EventDao)
	return ret0
}

func (_mr *_MockManagerRecorder) ServiceEventDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceEventDaoTransactions", arg0)
}

func (_m *MockManager) VersionInfoDao() dao.VersionInfoDao {
	ret := _m.ctrl.Call(_m, "VersionInfoDao")
	ret0, _ := ret[0].(dao.VersionInfoDao)
	return ret0
}

func (_mr *_MockManagerRecorder) VersionInfoDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VersionInfoDao")
}

func (_m *MockManager) VersionInfoDaoTransactions(db *gorm.DB) dao.VersionInfoDao {
	ret := _m.ctrl.Call(_m, "VersionInfoDaoTransactions", db)
	ret0, _ := ret[0].(dao.VersionInfoDao)
	return ret0
}

func (_mr *_MockManagerRecorder) VersionInfoDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VersionInfoDaoTransactions", arg0)
}

func (_m *MockManager) RegionUserInfoDao() dao.RegionUserInfoDao {
	ret := _m.ctrl.Call(_m, "RegionUserInfoDao")
	ret0, _ := ret[0].(dao.RegionUserInfoDao)
	return ret0
}

func (_mr *_MockManagerRecorder) RegionUserInfoDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegionUserInfoDao")
}

func (_m *MockManager) RegionUserInfoDaoTransactions(db *gorm.DB) dao.RegionUserInfoDao {
	ret := _m.ctrl.Call(_m, "RegionUserInfoDaoTransactions", db)
	ret0, _ := ret[0].(dao.RegionUserInfoDao)
	return ret0
}

func (_mr *_MockManagerRecorder) RegionUserInfoDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegionUserInfoDaoTransactions", arg0)
}

func (_m *MockManager) RegionAPIClassDao() dao.RegionAPIClassDao {
	ret := _m.ctrl.Call(_m, "RegionAPIClassDao")
	ret0, _ := ret[0].(dao.RegionAPIClassDao)
	return ret0
}

func (_mr *_MockManagerRecorder) RegionAPIClassDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegionAPIClassDao")
}

func (_m *MockManager) RegionAPIClassDaoTransactions(db *gorm.DB) dao.RegionAPIClassDao {
	ret := _m.ctrl.Call(_m, "RegionAPIClassDaoTransactions", db)
	ret0, _ := ret[0].(dao.RegionAPIClassDao)
	return ret0
}

func (_mr *_MockManagerRecorder) RegionAPIClassDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegionAPIClassDaoTransactions", arg0)
}

func (_m *MockManager) RegionProcotolsDao() dao.RegionProcotolsDao {
	ret := _m.ctrl.Call(_m, "RegionProcotolsDao")
	ret0, _ := ret[0].(dao.RegionProcotolsDao)
	return ret0
}

func (_mr *_MockManagerRecorder) RegionProcotolsDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegionProcotolsDao")
}

func (_m *MockManager) RegionProcotolsDaoTransactions(db *gorm.DB) dao.RegionProcotolsDao {
	ret := _m.ctrl.Call(_m, "RegionProcotolsDaoTransactions", db)
	ret0, _ := ret[0].(dao.RegionProcotolsDao)
	return ret0
}

func (_mr *_MockManagerRecorder) RegionProcotolsDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegionProcotolsDaoTransactions", arg0)
}

func (_m *MockManager) NotificationEventDao() dao.NotificationEventDao {
	ret := _m.ctrl.Call(_m, "NotificationEventDao")
	ret0, _ := ret[0].(dao.NotificationEventDao)
	return ret0
}

func (_mr *_MockManagerRecorder) NotificationEventDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NotificationEventDao")
}

func (_m *MockManager) AppBackupDao() dao.AppBackupDao {
	ret := _m.ctrl.Call(_m, "AppBackupDao")
	ret0, _ := ret[0].(dao.AppBackupDao)
	return ret0
}

func (_mr *_MockManagerRecorder) AppBackupDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AppBackupDao")
}

func (_m *MockManager) ServiceSourceDao() dao.ServiceSourceDao {
	ret := _m.ctrl.Call(_m, "ServiceSourceDao")
	ret0, _ := ret[0].(dao.ServiceSourceDao)
	return ret0
}

func (_mr *_MockManagerRecorder) ServiceSourceDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceSourceDao")
}

func (_m *MockManager) CertificateDao() dao.CertificateDao {
	ret := _m.ctrl.Call(_m, "CertificateDao")
	ret0, _ := ret[0].(dao.CertificateDao)
	return ret0
}

func (_mr *_MockManagerRecorder) CertificateDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CertificateDao")
}

func (_m *MockManager) CertificateDaoTransactions(db *gorm.DB) dao.CertificateDao {
	ret := _m.ctrl.Call(_m, "CertificateDaoTransactions", db)
	ret0, _ := ret[0].(dao.CertificateDao)
	return ret0
}

func (_mr *_MockManagerRecorder) CertificateDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CertificateDaoTransactions", arg0)
}

func (_m *MockManager) RuleExtensionDao() dao.RuleExtensionDao {
	ret := _m.ctrl.Call(_m, "RuleExtensionDao")
	ret0, _ := ret[0].(dao.RuleExtensionDao)
	return ret0
}

func (_mr *_MockManagerRecorder) RuleExtensionDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RuleExtensionDao")
}

func (_m *MockManager) RuleExtensionDaoTransactions(db *gorm.DB) dao.RuleExtensionDao {
	ret := _m.ctrl.Call(_m, "RuleExtensionDaoTransactions", db)
	ret0, _ := ret[0].(dao.RuleExtensionDao)
	return ret0
}

func (_mr *_MockManagerRecorder) RuleExtensionDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RuleExtensionDaoTransactions", arg0)
}

func (_m *MockManager) HttpRuleDao() dao.HttpRuleDao {
	ret := _m.ctrl.Call(_m, "HttpRuleDao")
	ret0, _ := ret[0].(dao.HttpRuleDao)
	return ret0
}

func (_mr *_MockManagerRecorder) HttpRuleDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HttpRuleDao")
}

func (_m *MockManager) HttpRuleDaoTransactions(db *gorm.DB) dao.HttpRuleDao {
	ret := _m.ctrl.Call(_m, "HttpRuleDaoTransactions", db)
	ret0, _ := ret[0].(dao.HttpRuleDao)
	return ret0
}

func (_mr *_MockManagerRecorder) HttpRuleDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HttpRuleDaoTransactions", arg0)
}

func (_m *MockManager) TcpRuleDao() dao.TcpRuleDao {
	ret := _m.ctrl.Call(_m, "TcpRuleDao")
	ret0, _ := ret[0].(dao.TcpRuleDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TcpRuleDao() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TcpRuleDao")
}

func (_m *MockManager) TcpRuleDaoTransactions(db *gorm.DB) dao.TcpRuleDao {
	ret := _m.ctrl.Call(_m, "TcpRuleDaoTransactions", db)
	ret0, _ := ret[0].(dao.TcpRuleDao)
	return ret0
}

func (_mr *_MockManagerRecorder) TcpRuleDaoTransactions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TcpRuleDaoTransactions", arg0)
}
