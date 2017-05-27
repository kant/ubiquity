package block_device_utils_test

import (
    "log"
    "github.com/IBM/ubiquity/remote/mounter/block_device_utils"
    "github.com/IBM/ubiquity/fakes"
    "github.com/IBM/ubiquity/utils"
    . "github.com/onsi/gomega"
    . "github.com/onsi/ginkgo"
    "testing"
    "errors"
    "fmt"
    "os"
    "io/ioutil"
)

var _ = Describe("block_device_utils_test", func() {
    var (
        logger        *log.Logger
        fakeExec      *fakes.FakeExecutor
        bdUtils       block_device_utils.BlockDeviceUtils
        err           error
        cmdErr        error = errors.New("command error")
    )

    BeforeEach(func() {
        logger = log.New(os.Stdout, "block_device_utils: ", log.Lshortfile|log.LstdFlags)
        fakeExec = new(fakes.FakeExecutor)
    })

    Context(".Rescan", func() {
        It("Rescan ISCSI calls 'sudo iscsiadm -m session --rescan'", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            err = bdUtils.Rescan()
            Expect(err).ToNot(HaveOccurred())
            Expect(fakeExec.ExecuteCallCount()).To(Equal(1))
            cmd, args := fakeExec.ExecuteArgsForCall(0)
            Expect(cmd).To(Equal("sudo"))
            Expect(args).To(Equal([]string{"iscsiadm", "-m", "session", "--rescan"}))
        })
        It("Rescan SCSI calls 'sudo rescan-scsi-bus -r'", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.SCSI, fakeExec)
            err = bdUtils.Rescan()
            Expect(err).ToNot(HaveOccurred())
            Expect(fakeExec.ExecuteCallCount()).To(Equal(1))
            cmd, args := fakeExec.ExecuteArgsForCall(0)
            Expect(cmd).To(Equal("sudo"))
            Expect(args).To(Equal([]string{"rescan-scsi-bus", "-r"}))
        })
        It("Rescan ISCSI fails if iscsiadm command missing", func() {
            fakeExec.IsExecutableReturns(cmdErr)
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            err = bdUtils.Rescan()
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
            Expect(fakeExec.ExecuteCallCount()).To(Equal(0))
            Expect(fakeExec.IsExecutableCallCount()).To(Equal(1))
            Expect(fakeExec.IsExecutableArgsForCall(0)).To(Equal("iscsiadm"))
        })
        It("Rescan SCSI fails if rescan-scsi-bus command missing", func() {
            fakeExec.IsExecutableReturns(cmdErr)
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.SCSI, fakeExec)
            err = bdUtils.Rescan()
            Expect(err).To(HaveOccurred())
            Expect(fakeExec.ExecuteCallCount()).To(Equal(0))
            Expect(fakeExec.IsExecutableCallCount()).To(Equal(2))
            Expect(fakeExec.IsExecutableArgsForCall(0)).To(Equal("rescan-scsi-bus"))
            Expect(fakeExec.IsExecutableArgsForCall(1)).To(Equal("rescan-scsi-bus.sh"))
        })
        It("Rescan ISCSI fails if iscsiadm execution fails", func() {
            fakeExec.ExecuteReturns([]byte{}, cmdErr)
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            err = bdUtils.Rescan()
            Expect(err).To(MatchError(cmdErr))
        })
        It("Rescan SCSI fails if rescan-scsi-bus execution fails", func() {
            fakeExec.ExecuteReturns([]byte{}, cmdErr)
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.SCSI, fakeExec)
            err = bdUtils.Rescan()
            Expect(err).To(MatchError(cmdErr))
        })
    })
    Context(".Discover", func() {
        It("Discover returns path for volume", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            volumeId := "volume-id"
            result := "mpath"
            fakeExec.ExecuteReturnsOnCall(1, []byte(
                fmt.Sprintf("%s (%s) dm-1 IBM", result, volumeId)), nil)
            mpath, err := bdUtils.Discover(volumeId)
            Expect(err).ToNot(HaveOccurred())
            Expect(mpath).To(Equal("/dev/mapper/" + result))
            Expect(fakeExec.ExecuteCallCount()).To(Equal(2))
            cmd1, args1 := fakeExec.ExecuteArgsForCall(0)
            Expect(cmd1).To(Equal("sudo"))
            Expect(args1).To(Equal([]string{"multipath"}))
            cmd2, args2 := fakeExec.ExecuteArgsForCall(1)
            Expect(cmd2).To(Equal("sudo"))
            Expect(args2).To(Equal([]string{"multipath", "-ll"}))
        })
        It("Discover fails if multipath command is missing", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            volumeId := "volume-id"
            fakeExec.IsExecutableReturns(cmdErr)
            _, err := bdUtils.Discover(volumeId)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
        It("Discover fails if multipath command fails", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            volumeId := "volume-id"
            fakeExec.ExecuteReturns([]byte{}, cmdErr)
            _, err := bdUtils.Discover(volumeId)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
        It("Discover fails if multipath -ll command fails", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            volumeId := "volume-id"
            fakeExec.ExecuteReturnsOnCall(1, []byte{}, cmdErr)
            _, err := bdUtils.Discover(volumeId)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
            Expect(fakeExec.ExecuteCallCount()).To(Equal(2))
        })
        It("Discover fails if volume not found", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            volumeId := "volume-id"
            fakeExec.ExecuteReturnsOnCall(1, []byte(
                fmt.Sprintf("mpath (other-volume) dm-1 IBM\nmpath (volume-id) dm-1 other-vendor")), nil)
            _, err := bdUtils.Discover(volumeId)
            Expect(err).To(HaveOccurred())
            Expect(fakeExec.ExecuteCallCount()).To(Equal(2))
        })
    })
    Context(".Cleanup", func() {
        It("Cleanup calls dmsetup and multipath", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            err = bdUtils.Cleanup(mpath)
            Expect(err).ToNot(HaveOccurred())
            Expect(fakeExec.ExecuteCallCount()).To(Equal(2))
            cmd1, args1 := fakeExec.ExecuteArgsForCall(0)
            Expect(cmd1).To(Equal("sudo"))
            Expect(args1).To(Equal([]string{"dmsetup", "message", mpath, "0", "fail_if_no_path"}))
            cmd2, args2 := fakeExec.ExecuteArgsForCall(1)
            Expect(cmd2).To(Equal("sudo"))
            Expect(args2).To(Equal([]string{"multipath", "-f", mpath}))
        })
        It("Cleanup fails if dmsetup command missing", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fakeExec.IsExecutableReturns(cmdErr)
            err = bdUtils.Cleanup(mpath)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
        It("Cleanup fails if dmsetup command fails", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fakeExec.ExecuteReturns([]byte{}, cmdErr)
            err = bdUtils.Cleanup(mpath)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
        It("Cleanup fails if multipath command missing", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "/dev/mapper/mpath"
            fakeExec.IsExecutableReturnsOnCall(1, cmdErr)
            err = bdUtils.Cleanup(mpath)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
            Expect(fakeExec.IsExecutableCallCount()).To(Equal(2))
        })
        It("Cleanup fails if multipath command fails", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fakeExec.ExecuteReturnsOnCall(1, []byte{}, cmdErr)
            err = bdUtils.Cleanup(mpath)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
    })
    Context(".CheckFs", func() {
        It("CheckFs detects exiting filesystem on device", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fakeExec.ExecuteReturns([]byte{}, nil)
            fs, err := bdUtils.CheckFs(mpath)
            Expect(err).ToNot(HaveOccurred())
            Expect(fs).To(Equal(false))
            Expect(fakeExec.ExecuteCallCount()).To(Equal(1))
            cmd, args := fakeExec.ExecuteArgsForCall(0)
            Expect(cmd).To(Equal("sudo"))
            Expect(args).To(Equal([]string{"blkid", mpath}))
        })
        It("CheckFs detects empty device", func() {
            err = ioutil.WriteFile("/tmp/tst.sh", []byte("exit 2"), 0777)
            Expect(err).ToNot(HaveOccurred())
            executor := utils.NewExecutor(logger)
            _, exitErr2 := executor.Execute("sh", []string{"/tmp/tst.sh"})
            Expect(exitErr2).To(HaveOccurred())
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fakeExec.ExecuteReturns([]byte{}, exitErr2)
            fs, err := bdUtils.CheckFs(mpath)
            Expect(err).ToNot(HaveOccurred())
            Expect(fs).To(Equal(true))
            Expect(fakeExec.ExecuteCallCount()).To(Equal(1))
            cmd, args := fakeExec.ExecuteArgsForCall(0)
            Expect(cmd).To(Equal("sudo"))
            Expect(args).To(Equal([]string{"blkid", mpath}))
        })
        It("CheckFs fails if blkid missing", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fakeExec.IsExecutableReturns(cmdErr)
            _, err = bdUtils.CheckFs(mpath)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
        It("CheckFs fails if blkid fails", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fakeExec.ExecuteReturns([]byte{}, cmdErr)
            _, err := bdUtils.CheckFs(mpath)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
    })
    Context(".MakeFs", func() {
        It("MakeFs creates fs by type", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fstype := "fstype"
            err = bdUtils.MakeFs(mpath, fstype)
            Expect(err).To(Not(HaveOccurred()))
            Expect(fakeExec.ExecuteCallCount()).To(Equal(1))
            cmd, args := fakeExec.ExecuteArgsForCall(0)
            Expect(cmd).To(Equal("sudo"))
            Expect(args).To(Equal([]string{"mkfs", "-t", fstype, mpath}))
        })
        It("MakeFs fails if mkfs missing", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fakeExec.IsExecutableReturns(cmdErr)
            err = bdUtils.MakeFs(mpath, "")
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
        It("MakeFs fails if mkfs command fails", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            fakeExec.ExecuteReturns([]byte{}, cmdErr)
            err = bdUtils.MakeFs(mpath, "")
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
    })
    Context(".MountFs", func() {
        It("MountFs succeeds", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            mpoint := "mpoint"
            err = bdUtils.MountFs(mpath, mpoint)
            Expect(err).To(Not(HaveOccurred()))
            Expect(fakeExec.ExecuteCallCount()).To(Equal(1))
            cmd, args := fakeExec.ExecuteArgsForCall(0)
            Expect(cmd).To(Equal("sudo"))
            Expect(args).To(Equal([]string{"mount", mpath, mpoint}))
        })
        It("MountFs fails if mount command missing", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            mpoint := "mpoint"
            fakeExec.IsExecutableReturns(cmdErr)
            err = bdUtils.MountFs(mpath, mpoint)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
        It("MountFs fails if mount command fails", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpath := "mpath"
            mpoint := "mpoint"
            fakeExec.ExecuteReturns([]byte{}, cmdErr)
            err = bdUtils.MountFs(mpath, mpoint)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })

    })
    Context(".UmountFs", func() {
        It("UmountFs succeeds", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpoint := "mpoint"
            err = bdUtils.UmountFs(mpoint)
            Expect(err).To(Not(HaveOccurred()))
            Expect(fakeExec.ExecuteCallCount()).To(Equal(1))
            cmd, args := fakeExec.ExecuteArgsForCall(0)
            Expect(cmd).To(Equal("sudo"))
            Expect(args).To(Equal([]string{"umount", mpoint}))
        })
        It("UmountFs fails if umount command missing", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpoint := "mpoint"
            fakeExec.IsExecutableReturns(cmdErr)
            err = bdUtils.UmountFs(mpoint)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
        It("UmountFs fails if umount command fails", func() {
            bdUtils = block_device_utils.GetBlockDeviceUtilsWithExecutor(logger, block_device_utils.ISCSI, fakeExec)
            mpoint := "mpoint"
            fakeExec.ExecuteReturns([]byte{}, cmdErr)
            err = bdUtils.UmountFs(mpoint)
            Expect(err).To(HaveOccurred())
            Expect(err).To(MatchError(cmdErr))
        })
    })
})


func TestGetBlockDeviceUtils(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "BlockDeviceUtils Test Suite")
}
